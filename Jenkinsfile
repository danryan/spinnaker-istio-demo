// https://issues.jenkins-ci.org/browse/JENKINS-39801
// import static java.util.UUID.randomUUID
def label = "demo-${UUID.randomUUID().toString()}"
def url = "https://github.com/danryan/spinnaker-istio-demo"
def namespace = "danryan"
def app = "spinnaker-demo"

podTemplate(
  serviceAccount: 'jenkins', 
  label: label, 
  containers: [
    containerTemplate(
      name: 'golang',
      image: 'golang:1.10',
      command: 'cat',
      ttyEnabled: true,
    ),
    containerTemplate(
      name: 'docker',
      image: 'docker:18.06.0-ce',
      command: 'cat',
      ttyEnabled: true,
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
  node(label) {
    stage("Checkout") {}
    // stage("Run Tests") {
    //   container('golang') {
    //     sh """
    //     echo 'Running tests!'
    //     """
    //   }
    // }

    stage("Build") {
      def myRepo = checkout scm
      def gitCommit = myRepo.GIT_COMMIT
      def gitBranch = myRepo.GIT_BRANCH
      def shortGitCommit = "${gitCommit[0..10]}"
      def previousGitCommit = sh(script: "git rev-parse ${gitCommit}~", returnStdout: true)

      def tag = (branch == "master") ? "latest" : branch

      container('docker') {
        git url: url, credentialsId: 'github:bco-jenkins-us-west-2'
        withCredentials([[$class: 'UsernamePasswordMultiBinding',
          credentialsId: 'danryan',
          usernameVariable: 'DOCKER_HUB_USER',
          passwordVariable: 'DOCKER_HUB_PASSWORD']]
        ){
          sh """
            docker login -u ${DOCKER_HUB_USER} -p ${DOCKER_HUB_PASSWORD}
            docker build -t ${namespace}/${app}:${tag} .
            docker tag ${namespace}/${app}:${tag} ${namespace}/${app}:${gitCommit} 
            docker push ${namespace}/${app}:${tag}
            docker push ${namespace}/${app}:${gitCommit}
          """
        }
      }
    }
  }
}
