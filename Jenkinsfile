podTemplate(containers: [
    containerTemplate(
        name: 'docker', 
        image: 'docker:latest',
        ttyEnabled: true,
        command: "cat")
        ],
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
            /* Ideally, we would run a test framework against our image.
            * For this example, we're using a Volkswagen-type approach ;-) */

            app.inside {
                sh 'echo "Tests passed"'
            }
        }

        stage('Push image') {
            /* Finally, we'll push the image with two tags:
            * First, the incremental build number from Jenkins
            * Second, the 'latest' tag.
            * Pushing multiple tags is cheap, as all the layers are reused. */
            docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-credentials') {
                app.push("${env.BUILD_NUMBER}")
                app.push("latest")
            }
        }
    }
}
