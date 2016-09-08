var path = require('path');
var gulp = require('gulp');
var conf = require('./conf');

var $ = require('gulp-load-plugins')({
    pattern: ['gulp-*', 'main-bower-files', 'del']
});

gulp.task('copy', function () {
    var fonts = ['bower_components/font-awesome/fonts/*'];
    var source = ['pics/**', 'conf.js', 'stack_samples/**', 'misc-tools/**'];

    gulp.src(fonts)
        .pipe(gulp.dest('dist/fonts'));
    gulp.src(source, {base: '.'})
        .pipe(gulp.dest('dist/'));
});

//directives html
gulp.task('template-min-directives', function () {
    return gulp.src('src/directives/**/*.html')
        .pipe($.minifyHtml({
            empty: true,
            spare: true,
            quotes: true
        }))
        .pipe($.angularTemplatecache('templateCacheHtmlDirectives.js', {
            module: 'app',
            root: '/src/directives'
        }))
        .pipe(gulp.dest('dist/src/'));
});

//utils html to js
gulp.task('template-min-utils', ['template-min-directives'], function () {
    return gulp.src('src/utils/**/*.html')
        .pipe($.minifyHtml({
            empty: true,
            spare: true,
            quotes: true
        }))
        .pipe($.angularTemplatecache('templateCacheHtmlUtils.js', {
            module: 'app.utils',
            root: '/src/utils'
        }))
        .pipe(gulp.dest('dist/src/'));
});

//auth html to js
gulp.task('template-min-auth', ['template-min-utils'], function () {
    return gulp.src('src/auth/**/*.html')
        .pipe($.minifyHtml({
            empty: true,
            spare: true,
            quotes: true
        }))
        .pipe($.angularTemplatecache('authTemplateCacheHtmlAuth.js', {
            module: 'app.auth',
            root: '/src/auth'
        }))
        .pipe(gulp.dest('dist/src/'));
});

gulp.task('template-min-layout', ['template-min-auth'], function () {
    return gulp.src('src/layout/**/*.html')
        .pipe($.minifyHtml({
            empty: true,
            spare: true,
            quotes: true
        }))
        .pipe($.angularTemplatecache('templateCacheHtmlLayout.js', {
            module: 'app.layout',
            root: '/src/layout'
        }))
        .pipe(gulp.dest('dist/src/'));
});

gulp.task('template-min-user', ['template-min-layout'], function () {
    return gulp.src('src/user/**/*.html')
        .pipe($.minifyHtml({
            empty: true,
            spare: true,
            quotes: true
        }))
        .pipe($.angularTemplatecache('templateCacheHtmlUser.js', {
            module: 'app.user',
            root: '/src/user'
        }))
        .pipe(gulp.dest('dist/src/'));
});

gulp.task('template-min-stack', ['template-min-user'], function () {
    return gulp.src('src/stack/**/*.html')
        .pipe($.minifyHtml({
            empty: true,
            spare: true,
            quotes: true
        }))
        .pipe($.angularTemplatecache('templateCacheHtmlStack.js', {
            module: 'app.stack',
            root: '/src/stack'
        }))
        .pipe(gulp.dest('dist/src/'));
});

gulp.task('template-min-node', ['template-min-stack'], function () {
    return gulp.src('src/node/**/*.html')
        .pipe($.minifyHtml({
            empty: true,
            spare: true,
            quotes: true
        }))
        .pipe($.angularTemplatecache('templateCacheHtmlNode.js', {
            module: 'app.node',
            root: '/src/node'
        }))
        .pipe(gulp.dest('dist/src/'));
});

gulp.task('template-min-network', ['template-min-node'], function () {
    return gulp.src('src/network/**/*.html')
        .pipe($.minifyHtml({
            empty: true,
            spare: true,
            quotes: true
        }))
        .pipe($.angularTemplatecache('templateCacheHtmlNetwork.js', {
            module: 'app.network',
            root: '/src/network'
        }))
        .pipe(gulp.dest('dist/src/'));
});

gulp.task('template-min-registry', ['template-min-network'], function () {
    return gulp.src('src/registry/**/*.html')
        .pipe($.minifyHtml({
            empty: true,
            spare: true,
            quotes: true
        }))
        .pipe($.angularTemplatecache('templateCacheHtmlRegistry.js', {
            module: 'app.registry',
            root: '/src/registry'
        }))
        .pipe(gulp.dest('dist/src/'));
});

gulp.task('template-min-registryAuth', ['template-min-registry'], function () {
    return gulp.src('src/registry-auth/**/*.html')
        .pipe($.minifyHtml({
            empty: true,
            spare: true,
            quotes: true
        }))
        .pipe($.angularTemplatecache('templateCacheHtmlReAuth.js', {
            module: 'app.registryAuth',
            root: '/src/registry-auth'
        }))
        .pipe(gulp.dest('dist/src/'));
});

gulp.task('template-min-misc', ['template-min-registryAuth'], function () {
    return gulp.src('src/misc/**/*.html')
        .pipe($.minifyHtml({
            empty: true,
            spare: true,
            quotes: true
        }))
        .pipe($.angularTemplatecache('templateCacheHtmlMisc.js', {
            module: 'app.misc',
            root: '/src/misc'
        }))
        .pipe(gulp.dest('dist/src/'));
});

gulp.task('ng-annotate', ['template-min-misc'], function () {
    return gulp.src('src/**/*.js')
        .pipe($.ngAnnotate({add: true}))
        .pipe(gulp.dest('dist/src/'))
});

gulp.task('html-replace', ['ng-annotate'], function () {
    var templateInjectFile = gulp.src('dist/src/templateCacheHtml*.js', {read: false});
    var templateInjectOptions = {
        starttag: '<!-- inject:template.js  -->',
        addRootSlash: false
    };

    var revAll = new $.revAll();
    return gulp.src('index.html')
        .pipe($.inject(templateInjectFile, templateInjectOptions))
        .pipe($.useref()).on('error', $.util.log)
        .pipe($.if('*.js', $.uglify()))
        .pipe($.if('*.css', $.minifyCss()))
        .pipe(revAll.revision().on('error', $.util.log))
        .pipe($.revHash())
        .pipe(gulp.dest('dist/'))
        .pipe(revAll.manifestFile())
        .pipe(gulp.dest('dist/'));
});

gulp.task('html-rename', ['html-replace'], function () {
    gulp.src('dist/index.*.html')
        .pipe($.rename('index.html').on('error', $.util.log))
        .pipe(gulp.dest('dist/'));
});

gulp.task('auth-html-replace', ['html-rename'], function () {

    var templateInjectFile = gulp.src('dist/src/authTemplateCacheHtml*.js', {read: false});
    var templateInjectOptions = {
        starttag: '<!-- inject:template.js  -->',
        addRootSlash: false
    };

    var revAll = new $.revAll();
    return gulp.src('auth-index.html')
        .pipe($.inject(templateInjectFile, templateInjectOptions))
        .pipe($.useref()).on('error', $.util.log)
        .pipe($.if('*.js', $.uglify()))
        .pipe($.if('*.css', $.minifyCss()))
        .pipe(revAll.revision().on('error', $.util.log))
        .pipe($.revHash())
        .pipe(gulp.dest('dist/'))
        .pipe(revAll.manifestFile())
        .pipe(gulp.dest('dist/'));
});

gulp.task('auth-html-rename', ['auth-html-replace'], function () {
    gulp.src('dist/auth-index.*.html')
        .pipe($.rename('auth-index.html').on('error', $.util.log))
        .pipe(gulp.dest('dist/'));
});

gulp.task('clean', ['auth-html-rename'], function () {
    return $.del([path.join(conf.paths.dist, '/src'), path.join(conf.paths.dist, 'auth-index.*.html'), path.join(conf.paths.dist, 'index.*.html')]);
});

gulp.task('delete', function () {
    return $.del([path.join(conf.paths.dist, '/'), path.join(conf.paths.tmp, '/')]);
});

gulp.task('build', ['copy', 'clean']);