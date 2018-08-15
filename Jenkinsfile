pipeline {
  agent {
    kubernetes {
      label 'mypod'
      containerTemplate {
        name 'curl'
        image 'tutum/curl'
        ttyEnabled true
        command 'cat'
      }
    }
  }
  environment {
    CONTAINER_ENV_VAR = 'container-env-var-value'
  }
  stages {
      steps {
        container('curl') {
          sh 'echo INSIDE_CONTAINER_ENV_VAR = ${CONTAINER_ENV_VAR}'
        }
        sh 'echo OUTSIDE_CONTAINER_ENV_VAR = ${CONTAINER_ENV_VAR}'
      }
    }
  }
}

