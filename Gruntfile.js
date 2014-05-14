module.exports = function(grunt) {
  // require('logfile-grunt')(grunt);

  // Project configuration.
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),

    go: {
      options: {
        GOPATH: ["./"]
      },

      vadar: {
        root: "src/github.com/vadar",
        output: "app",
        run_files: ["vadar.go"]
      }
    },

    watch: {
      files: ['**/*.go'],
      tasks: ['devBuild']
    }
  });

  grunt.loadNpmTasks('grunt-go');
  grunt.loadNpmTasks('grunt-contrib-watch');

  // Default task(s).
  grunt.registerTask('default', ['devBuild']);
  grunt.registerTask('devBuild', ['go:install:vadar', 'go:test:vadar']);
  grunt.registerTask('format', ['go:fmt:vadar']);
};
