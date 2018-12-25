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
                docker {
                    image 'obraun/vss-jenkins'
                    args '-v ${PWD}:/go/src/github.com/ob-vss-ws18/ob-vss-ws18/'
                }
            }
            steps {
                sh 'cd proto.actor/node1 && make app'
                sh 'cd proto.actor/node2 && make app'
            }
        }
        stage('Build Docker-Images') {
            agent any
            steps {
                sh 'cd proto.actor && docker-compose build'
            }
        }
        stage('Run Docker-Images') {
            agent any
            steps {
                sh 'cd proto.actor && docker-compose up -d'
                sleep(time:2,unit:"MINUTES")
                sh 'cd proto.actor && docker-compose down'
            }
        }
    }
}