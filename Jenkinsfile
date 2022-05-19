pipeline {
  agent any
  stages {
    stage('Start Build') {
      steps {
        echo 'Start Build'
        sh '''export GO111MODULE=auto
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