// https://issues.jenkins-ci.org/browse/JENKINS-39801
// import static java.util.UUID.randomUUID
def label = "worker-${UUID.randomUUID().toString()}"

podTemplate(
  serviceAccount: 'jenkins', 
  label: label, 
  containers: [
    containerTemplate(
      name: 'demo',
      image: 'danryan/spinnaker-istio-demo',
      command: 'true',
      ttyEnabled: true,
      resourceRequestCpu: '100m',
      resourceLimitMemory: '128Mi'
    // ),
    // containerTemplate(
    //   name: 'curl',
    //   image: 'tutum/curl',
    //   command: 'true',
    //   ttyEnabled: true,
    //   resourceRequestCpu: '100m',
    //   resourceLimitMemory: '128Mi'
    )
  ], 
  envVars: [
    envVar(key: 'BRANCH_NAME', value: env.BRANCH_NAME),
    envVar(key: 'BUILD_NUMBER', value: env.BUILD_NUMBER)
  ], 
  volumes: [
    hostPathVolume(mountPath: '/var/run/docker.sock', hostPath: '/var/run/docker.sock')
  ]
) 
{
  node("node-${label}") {
    def myRepo = checkout scm
    def gitCommit = myRepo.GIT_COMMIT
    def gitBranch = myRepo.GIT_BRANCH
    def shortGitCommit = "${gitCommit[0..10]}"
    def previousGitCommit = sh(script: "git rev-parse ${gitCommit}~", returnStdout: true)

    // stage('Checkout') {
    //   checkout scm
    // }
    stage('Test') {
      try {
        // container('demo') {
          sh 'echo "testing"'
        // }
      } finally {}
    }
    // stage('Package') {
    //   try {
    //     // container('demo') {
    //       sh 'echo "packaging"'
    //     // }
    //   } finally { }
    // }
    // stage('Deploy') {
    //   try {
    //     container('curl') {
    //       sh "curl -X POST -H 'Content-Type: application/json' -d '{\"build_url\":\"${env.BUILD_URL}\"}' https://api.spinnaker.k8s.us-east-2.bco.aws.cudaops.com/webhooks/webhook/demo-jenkins"
    //     }
    //   } finally {}
    // }
  }
}
