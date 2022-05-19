pipeline {
  agent any
  stages {
    stage('Start Build') {
      steps {
        echo 'Start Build'
        sh '''go mod init
go build'''
      }
    }

    stage('Test') {
      steps {
        echo 'Testing'
      }
    }

  }
}