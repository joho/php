class Quote < Struct.new(:text, :author, :link)
  def self.random
    all.shuffle[0]
  end

private
  def self.all
    [
      new("There is nothing to writing. All you do is sit down at a typewriter and bleed.", "Ernest Hemingway", nil),
      new("Don't you know who I am?", "Ben Schwarz", "http://germanforhipster.com"),
      new("I think we underestimate the complexity of deep sea cabling", "@deepjohnbarton", "https://twitter.com/deepjohnbarton/status/250782278349897728"),
      new("Ninety percent of everything is crap.", "Sturgeon's Law", "http://en.wikipedia.org/wiki/Sturgeon's_Law")
    ]
  end
end
