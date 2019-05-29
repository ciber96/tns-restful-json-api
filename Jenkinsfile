node {
   def root = tool name: '1.12', type: 'go'
   stage('Preparation') {
      // Get some code from a GitHub repository
      git branch: 'dev', url: 'https://github.com/ciber96/tns-restful-json-api'
      stash 'repo'
      // Export environment variables pointing to the directory where Go was installed
      withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
          sh label: '', script: 'go get github.com/gorilla/mux'
      }
   }
   stage('Unit Tests') {
      // Run the go tests
      withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        sh label: '', script: 'cd v2 && go test .'
      }
   }
   stage('Build') {
      // Run the go build
      withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        sh label: '', script: 'cd v2 && go build -o dev'
      }
      stash 'repo'
   }
   stage('Deploy') {
      sshPublisher(publishers: [sshPublisherDesc(configName: 'API Dev', transfers: [sshTransfer(cleanRemote: false, excludes: '', execCommand: 'if pgrep screen; then sudo pkill screen; fi', execTimeout: 120000, flatten: false, makeEmptyDirs: false, noDefaultExcludes: false, patternSeparator: '[, ]+', remoteDirectory: '', remoteDirectorySDF: false, removePrefix: '', sourceFiles: ''), sshTransfer(cleanRemote: false, excludes: '', execCommand: 'sudo chmod a+x \'/home/vagrant/dev\'', execTimeout: 120000, flatten: false, makeEmptyDirs: false, noDefaultExcludes: false, patternSeparator: '[, ]+', remoteDirectory: '', remoteDirectorySDF: false, removePrefix: 'v2/', sourceFiles: 'v2/dev'), sshTransfer(cleanRemote: false, excludes: '', execCommand: 'cd /home/vagrant/ && screen -dmS api ./dev', execTimeout: 120000, flatten: false, makeEmptyDirs: false, noDefaultExcludes: false, patternSeparator: '[, ]+', remoteDirectory: '', remoteDirectorySDF: false, removePrefix: '', sourceFiles: '')], usePromotionTimestamp: false, useWorkspaceInPromotion: false, verbose: false)])
   }
}