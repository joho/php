I want to share one of my favourite TypeScript techniques for mocking dependencies in tests.

My affection for this technique largely comes from an aesthetic developed as a Ruby and Go programmer, but TypeScript has a helper type that puts a really nice and neat ribbon on top.

https://medium.com/@cep21/what-accept-interfaces-return-structs-means-in-go-2fe879e25ee8

Let's take the following grossly simplified code as a starting point. We have a very vague domain entity called `Thing` and the associated repo like this:

```typescript
// entities/things.ts

export interface Thing {
    id: string
    isActive: boolean
}

export default interface ThingRepo {
    all() => Promise<Thing[]>
    find(id: string) => Promise<Thing>
    save(thing: Thing) => Promise<void>
}
```

We then write somewhere else some client code operating on things.

```typescript
// services/things.ts

import { Thing, ThingRepo } from 'entities/things'

export async function findAllActiveThings(repo: ThingRepo): Promise<Thing[]> {
    const repo = new ThingRepo()

    const things = await repo.all()

    const activeThings = things.filter((thing) => thing.isActive)

    return activeThings
}
```

Then finally a test that looks a bit like this. I just sketched this out directly in the post and haven't tried to run it, so forgive any minor errors.

```typescript
// services/things.test.ts

import { findAllActiveThings } from './things'

jest.mock('entities/things', () => {
  return {
        ThingRepo: jest.fn().mockImplementation(() => {
            return { all: () => {
                return Promise.resolve([
                    { id: '1', isActive: false},
                    { id: '2', isActive: true},
                ])
            }};
        });
    }
});

it('findAllActiveThings only returns active things', async () => {
    const things = await findAllActiveThings()
    expect(things.length).toBe(1)
    expect(things[0].isActive).toBe(true)
})
```
