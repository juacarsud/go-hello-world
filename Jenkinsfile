podTemplate(containers: [
    containerTemplate(
        name: 'docker', 
        image: 'docker:latest',
        ttyEnabled: true,
        command: "cat"),
    containerTemplate(
        name: 'golang',
        image: 'golang:latest',
        ttyEnabled: true,
        command: "cat"
        )],
        volumes: [
            hostPathVolume( mountPath: '/var/run/docker.sock', hostPath: '/var/run/docker.sock')
        ]
        ){
      node(POD_LABEL) {
        def app

        stage('Clone repository') {
        /* Let's make sure we have the repository cloned to our workspace */
            checkout scm
        }
        stage('Test image') {
            container('golang'){
                stage('Testing Code'){
                    sh 'cd ${GOPATH}/src'
                    sh 'mkdir -p ${GOPATH}/src/go-hello-world'
                    sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/go-hello-world'
                    sh 'go clean -cache'
                    sh 'go mod init hello'
                    sh 'go test -v -short'
                }
            }
        }
        stage('Build code') {
            container('golang'){
                stage('Build Code'){
                    sh 'cd ${GOPATH}/src'
                    sh 'mkdir -p ${GOPATH}/src/go-hello-world'
                    sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/go-hello-world'
                    sh 'go clean -cache'
                    sh 'go build'
                }
            }
        }
        stage('Build and Push image') {
            container('docker'){
                stage('Publish Docker Image'){
                    docker.withRegistry('https://073278647946.dkr.ecr.us-east-2.amazonaws.com', 'ecr:us-east-2:aws_credentials') {
                    app = docker.build("go-hello-app:1.0.${env.BUILD_NUMBER}")
                    app.push()
                    app.push('latest')
                    }
                }
            }
        }
    }
}
