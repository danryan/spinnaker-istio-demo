/**
 * Build and test the kubernetes plugin using the plugin itself in a Kubernetes cluster
 *
 * A `jenkins` service account needs to be created using src/main/kubernetes/service-account.yml
 *
 * A PersistentVolumeClaim needs to be created ahead of time with the definition in examples/maven-with-cache-pvc.yml
 *
 * NOTE that typically writable volumes can only be attached to one Pod at a time, so you can't execute
 * two concurrent jobs with this pipeline. Or change readOnly: true after the first run
 */

def label = UUID.randomUUID().toString()

// timestamps {
  podTemplate(
    serviceAccount: 'jenkins', 
    label: label, 
    containers: [
      containerTemplate(
        name: 'demo', 
        image: 'danryan/spinnaker-istio-demo', 
        ttyEnabled: true,
        resourceRequestCpu: '100m',
        resourceLimitMemory: '1200Mi'
      )
    ], 
    envVars: [
      envVar(key: 'BRANCH_NAME', value: env.BRANCH_NAME),
      envVar(key: 'BUILD_NUMBER', value: env.BUILD_NUMBER)
    ], 
    volumes: []
  ) 
  {
    node(label) {
      stage('Checkout') {
        checkout scm
      }
      stage('Test') {
        try {
          container('demo') {
            sh 'echo "testing"'
          }
        } finally {}
      }
      stage('Package') {
        try {
          container('demo') {
            sh 'echo "packaging"'
          }
        } finally { }
      }
      stage('Deploy') {
        try {
          sh "curl -X POST -H 'Content-Type: application/json' -d '{\"build_url\":\"${env.BUILD_URL}\"}' https://api.spinnaker.k8s.us-east-2.bco.aws.cudaops.com/webhooks/webhook/demo-jenkins"
        } finally {}
      }
    }
  }
// }
