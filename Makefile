# A simple Makefile alternative to using Grunt for your static asset compilation
#
## Usage
#
#   $ npm install
#
# And then you can run various commands:
#
#   $ make            # compile files that need compiling
#   $ make clean all  # remove target files and recompile from scratch
#   $ make watch      # watch the filesystem for changes and recompile
#
## Why?
#
# 1. Less dependencies
#
# Instead of needing a grunt-* wrapper library for your favourite tool, and
# needing to update or wait for a new version just to pass a new option into
# the underlying library, you can simply use the library directly.
#
# Thankfully all the awesome node libraries you're using with Grunt include
# command-line tools which are easily executable by make.
#
# 2. Easy to extend
#
# For this example we're using some linux commands for concating, and various
# node-based libraries for Sass compilation, CSS prefixing, etc.
#
# Adding a new tool or step to your asset compilation is dead simple:
#
#   1. add the library to package.json
#   2. npm install
#   3. add a new line to the revelant target, calling the binary created by npm
#      install
#
# 3. Performance
#
# Makefile understands file modification times, so it won't recompile any
# targets whose source dependencies haven't changed. Combined with using a
# file modification monitoring tool like wach, you get near-instant recompiles
# of your front-end assets.
 
 
# Variables
 
APP=myapp
APP_JS_SOURCES=$(BOWER)/angular/angular.js $(JS)/$(APP).js
AUTOPREFIXER_BROWSERS="> 1%"
 
SASS=views/stylesheets
 
BIN=node_modules/.bin
BOWER=bower_components
DIST=public
 
# Targets
#
# The format goes:
#
#   target: list of dependencies
#     commands to build target
#
# If something isn't re-compiling double-check the changed file is in the
# target's dependencies list.
 
# Compile the final targets
all: $(DIST)/styles.css $(DIST)/narrow.css $(DIST)/wide.css 
 
# Destroy the final targets
clean:
	rm -f $(DIST)/*.css
 
# Watch the filesystem and recompile on file modification
watch:
	$(BIN)/wach -o "$(SASS)/**/*" make all
 
# The final CSS file
$(DIST)/styles.css: $(SASS)/styles.scss
	$(BIN)/node-sass $(SASS)/styles.scss $(DIST)/styles.css
	$(BIN)/autoprefixer --browsers $(AUTOPREFIXER_BROWSERS) $(DIST)/styles.css
 
$(DIST)/wide.css: $(SASS)/wide.scss
	$(BIN)/node-sass $(SASS)/wide.scss $(DIST)/wide.css
	$(BIN)/autoprefixer --browsers $(AUTOPREFIXER_BROWSERS) $(DIST)/wide.css

$(DIST)/narrow.css: $(SASS)/narrow.scss
	$(BIN)/node-sass $(SASS)/narrow.scss $(DIST)/narrow.css
	$(BIN)/autoprefixer --browsers $(AUTOPREFIXER_BROWSERS) $(DIST)/narrow.css

