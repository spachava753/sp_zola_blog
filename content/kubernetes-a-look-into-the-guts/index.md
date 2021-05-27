+++
title = "Kubernetes: A look into the guts"
date = 2019-09-06
+++

{{ page.assets }}

**Kubernetes.** A word that is well-known and hyped throughout the industry. I won't go into all things great about Kubernetes, but here are just a few.

# Flexible and extensible

Kubernetes is designed to flexible and extensible. You can see the an example of this by looking at the Load Balancing resource. Each Kubernetes platform provider provisions this resource a different way, but the actual behavior does not change. The load balancer could be provisioned by AWS, GCP, or just plain Minikube. Regardless, your application will work on all of these platforms. Another example is all the overlay networks available, such as Calico, Weave, and Flannel. Kubernetes can work with any one of these. However, Kubernetes is also extensible. Kubernetes allows you to define your own resources and controllers. Calico is a great example of this, as it allows you to define custom policies and global rules on the network, which is not supported on vanilla Kubernetes.

# Cheap and Scalable

If you need to explain to a non-technical person why Kubernetes is so great (like a manager), then talking about cost savings is your best bet. With Kubernetes, you don't have to worry about scaling down or up, as Kubernetes will handle that. Not only that, but Kubernetes has a cloud mode that allows you to run applications on hosts until hardware resources can no longer support more load. Especially with the cloud, companies have an easy way to scale out to meet any demand, from zero to Black Friday blowouts.

# The Control Plane

The control plane of Kubernetes is where the orchestration magic happens. It keeps records of all the resources present in the cluster, what kind of resources are available, and the state of applications in cluster. The website gives a proper description:

> The various parts of the Kubernetes Control Plane, such as the Kubernetes Master and kubelet processes, govern how Kubernetes communicates with your cluster. The Control Plane maintains a record of all of the Kubernetes Objects in the system, and runs continuous control loops to manage those objects’ state. At any given time, the Control Plane’s control loops will respond to changes in the cluster and work to make the actual state of all the objects in the system match the desired state that you provided.

## But...What actually happens?

So now we now that the control pane is incredibly important, and is important to all of the processes running in your cluster. But we still don't know exactly how control pane works. How does it even provision pods to run on nodes? Who decides what to run where? This post is to explain that.

## The Process

{{ resize_image(path="kubernetes-a-look-into-the-guts/image_preview.png", width=800, height=800, op="fit") }}

First, we have a cluster that is set up, as well as a client with properly configured kubectl. We have two nodes connected to control pane.

{{ resize_image(path="kubernetes-a-look-into-the-guts/image_preview-1-.png", width=800, height=800, op="fit") }}

We run a kubectl command to communicate to the one of the key components of the control pane: the API server. The command communicates to the API server and tells it that we want three nginx instances running. You'll notice that it does not specify how. This is the declarative approach taken by Kubernetes, so that you only define the desired state, and Kubernetes will find a way to meet those expectations.

{{ resize_image(path="kubernetes-a-look-into-the-guts/image_preview-2-.png", width=800, height=800, op="fit") }}

The API server then creates a new resource in the etcd cluster called a Deployment. Deployments are in charge of creating ReplicaSets and managing Pod state. The reason why we don't create a ReplicaSet directly is because we would have to track and update Pod state ourselves given a Pod template update or patch. Deployment will take care of automatically patching in new Pod templates. That is just one of the use cases, there are many more.

{{ resize_image(path="kubernetes-a-look-into-the-guts/image_preview-3-.png", width=800, height=800, op="fit") }}

The API server sends a response back to the client stating that the Deployment resource has been created and is now being provisioned.

{{ resize_image(path="kubernetes-a-look-into-the-guts/image_preview-6-.png", width=800, height=800, op="fit") }}

The API server sends a signal to the controller manager, and the controller manager sees that a new Deployment resource has been created. It then proceeds to create another resource called the ReplicaSet, which is in charge of replicating Pods and managing Pods. If there is not enough Pods, it provisions some more, if there is too much, it removes some Pods.

{{ resize_image(path="kubernetes-a-look-into-the-guts/image_preview-9-.png", width=800, height=800, op="fit") }}

Once again the controller manager is notified and sees the ReplicaSet that it created, it creates the appropriate number of Pods to be provisioned on the nodes.

{{ resize_image(path="kubernetes-a-look-into-the-guts/image_preview-12-.png", width=800, height=800, op="fit") }}

Once the controller manager creates the pods, then the Scheduler is notified to decide where to put these Pods. The Scheduler picks the nodes based on specific policies, and the type of resource. For example, it could provision one Pod per node for DameonSets.

{{ resize_image(path="kubernetes-a-look-into-the-guts/image_preview-14-.png", width=800, height=800, op="fit") }}

The API server notifies the Kublet processes in each node to create the specified Pods based on what node the Scheduler set it to.

{{ resize_image(path="kubernetes-a-look-into-the-guts/image_preview-16-.png", width=800, height=800, op="fit") }}

The Kublet processes send an update to API server notifying it that the Pods are now running and ready.

# Conclusion

Hopefully, it's a little more clear now how the control pane provisions new resources. Of course, the specifics differ from resource to resource, but this is the general flow. The API server takes care of creating the resources in the etcd, and offloads the decisions on how to provision Pods to controller managers and the scheduler.

**Courtesy of container.training for the images.**