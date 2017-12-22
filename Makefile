 
AUTOPREFIXER_BROWSERS="> 1%"
SASS=views/stylesheets
BIN=node_modules/.bin
DIST=public
STYLESHEETS=$(subst ${SASS},${DIST}/stylesheets,$(patsubst %.scss,%.css,$(wildcard $(SASS)/*.scss)))
 
# Destroy the final targets
clean:
	rm -f $(DIST)/stylesheets/*.css
	rm -f web

# Compile the final targets
all: $(STYLESHEETS)

# Watch the filesystem and recompile on file modification
watch: 
	$(BIN)/watch "make all" views
 
# The final CSS files
$(DIST)/stylesheets/%.css: $(SASS)/%.scss
	$(BIN)/node-sass $< $@ 
	$(BIN)/autoprefixer-cli --browsers $(AUTOPREFIXER_BROWSERS) $@

