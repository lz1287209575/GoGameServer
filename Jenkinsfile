pipeline {
  agent any
  stages {
    stage('Start Build') {
      steps {
        echo 'Start Build'
        sh 'go build'
      }
    }

    stage('Test') {
      steps {
        echo 'Testing'
      }
    }

  }
}