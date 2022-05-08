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

        stage('Build image') {
            container('docker'){
                stage('Inside Container'){
                    sh '''
                    docker build -t test:latest .
                    '''
                }
            }
            /* This builds the actual image; synonymous to
            * docker build on the command line */
        }

        stage('Test image') {
            container('golang'){
                stage('Testing Code'){
                    sh 'cd ${GOPATH}/src'
                    sh 'mkdir -p ${GOPATH}/src/go-hello-world'
                    sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/go-hello-world'
                    sh 'go clean -cache'
                    sh 'ls -la'
                    sh 'go test -v -short'
                }
            }
        }

        stage('Push image') {
            container('docker'){
                stage('Publish Docker Image'){
                    docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-credentials') {
                    
            }
                }
            }
            /* Finally, we'll push the image with two tags:
            * First, the incremental build number from Jenkins
            * Second, the 'latest' tag.
            * Pushing multiple tags is cheap, as all the layers are reused. */
        }
    }
}
