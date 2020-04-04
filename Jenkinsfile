pipeline
    {
        environment
        {
            registry = credentials("go_docker_registry")
            registryCredential = 'dockerhub'
        }
        agent any
        stages
        {
            stage('Git Checkout')
 			{
 		        steps
 		   		{
 					checkout scm
 				}
 		    }
            stage('Build Docker Image')
			{
				steps
				{
					script
					{
						dockerImage = docker.build("${registry}:latest")
					}
				}
			}
			stage('push docker image')
			{
			    steps
			    {
			        script
			        {
			            docker.withRegistry( '', registryCredential )
							{
								dockerImage.push()
							}
			        }
			    }
			}
        }
    }