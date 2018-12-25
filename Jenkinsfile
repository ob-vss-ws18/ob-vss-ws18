pipeline {
    agent none
    stages {
        stage('Test') {
            agent {
                docker { image 'obraun/vss-jenkins' }
            }
            steps {
                sh 'echo no tests'
            }
        }
        stage('Lint') {
            agent {
                docker { image 'obraun/vss-jenkins' }
            }   
            steps {
                sh 'echo golangci-lint run --enable-all'
            }
        }
        stage('Build Nodes') {
            agent {
                docker { image 'obraun/vss-jenkins' }
            }
            steps {
                sh 'mkdir -p /go/src/github.com/ob-vss-ws18'
                sh 'ln -s $(pwd) /go/src/github.com/ob-vss-ws18/ob-vss-ws18'
                sh 'cd proto.actor/node1 && make app'
                sh 'cd proto.actor/node2 && make app'
            }
        }
        stage('Build and Push Docker-Images') {
            agent {
                label 'master'
            }
            steps {
                sh 'cd proto.actor/node1 && docker-build-and-push -b ${BRANCH_NAME} -s node1'
                sh 'cd proto.actor/node2 && docker-build-and-push -b ${BRANCH_NAME} -s node2'
            }
        }
    }
}