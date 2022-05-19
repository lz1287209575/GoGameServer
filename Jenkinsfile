pipeline {
  agent any
  stages {
    stage('Start Build') {
      steps {
        echo 'Start Build'
        sh '''
go build -o GoGameServer'''
      }
    }

    stage('Test') {
      steps {
        echo 'Testing'
      }
    }

  }
}