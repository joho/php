class ExternalLink < Struct.new(:title, :url)
  def self.recent_writings
    [
      new("Web Scale for the Rest of Us (RubyConfAU Talk)", "http://vimeo.com/61342269"),
      new("Interview for The Fetch", "http://blog.thefetch.com/2012/10/29/interview-melbourne-local-john-barton/"),
      new("The Minimum Viable Rails Stack", "http://goodfil.ms/blog/posts/2012/09/15/minimum-viable-rails-stack/"),
      new("Will App.net be the connective tissue founders & developers can rely on?", "http://thenextweb.com/apps/2012/08/18/why-10000-people-care-app-net/"),
      new("Scaling Rails @ Melbourne Ruby Meetup", "http://jrb.tumblr.com/post/30570014929/scaling-rails-at-melbourne-roro")
    ]
  end

  def self.interesting_projects
    [
      new("This site (Sinatra + SCSS)", "https://github.com/joho/whoisjohnbarton.com"),
      new("Spoilerless Tour De France", "https://github.com/joho/letour"),
      new("Loose (and fast) Postgres Row Counts", "https://github.com/goodfilms/postgres_loose_table_counts"),
      new("Netflix API wrapper for Ruby", "https://github.com/goodfilms/netflix-ruby"),
      new("HTTP Status codes 7xx (Developer Fouls)", "https://github.com/joho/7XX-rfc")
    ]
  end

  def self.social_networks
    [
      new("twitter.com/johnbarton", "http://twitter.com/johnbarton"),
      new("goodfil.ms/john", "http://goodfil.ms/john"),
      new("github.com/joho", "http://github.com/joho"),
      new("jrb.tumblr.com", "http://jrb.tumblr.com"),
      new("goodreads.com/johnbarton", "http://goodreads.com/johnbarton"),
      new("au.linkedin.com/in/johnbarton", "http://au.linkedin.com/in/johnbarton")
    ]
  end
end
