def build_docker_image() {
    return {
        stage("Building image"){
            echo "Build Docker Image devopsabi/ab_demo_app:${new_version}"            
            script{
                withDockerRegistry(url:'https://index.docker.io/v1/',credentialsId:'docker_hub_id'){
                    current_deploy_version = sh (returnStdout: true, script: 'cat new_version.txt').trim()
                    docker.build('devopsabi/ab_demo_app:'+current_deploy_version).push()
                }
            }
            
        }
    }
}

def update_app() {
    return {
        stage("Updating app to ${new_version}")
        new_version = sh(script: 'cat new_version.txt', returnStdout: true)
        deploy_app = sh(script: 'docker service update --image devopsabi/ab_demo_app`cat new_version.txt` demo_app', returnStdout: true)
        echo "${deploy_app}"
    }
}

 pipeline {
        agent any
        
        stages {
        
            stage('checkout_demo_app') {
                steps {
                    git branch:"main", url: 'https://github.com/devopsabi/Docker_Test.git', credentialsId:'githubab'
                }
    
            }

         stage('build docker image demo_app'){
                steps {
                    script{

			    whoami_user = sh(script: 'whoami')
        		    echo "${whoami_user}"
                            sh label: 'Import GPG key', script: 'HOME=$(pwd) gpg --import AB-GPG-KEY'
                            status = sh(script: 'cat deploy.txt', returnStdout: true).trim()
                            current_version = sh(script: 'cat current_version.txt', returnStdout: true)
                            new_version = sh(script: 'cat new_version.txt', returnStdout: true)
                            println "${status}"
                            if ( status == 'yes' ){
                                    println "current Version = ${current_version}\nNew version = ${new_version}"
                                    build_docker_image().call()
                                }
                                
                            else {
                                    println "NOT DEPLOYING THE CHANGES... Please update deploy.txt and set as 'yes' If you would like to deploy the current version "

                                }
                    
                    }
                }
            }
            

    }
}
