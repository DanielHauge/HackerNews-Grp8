# LSD Exam Report
- By Group L: Emmely (cph-el69) & Kristian (cph-kf96) & Daniel (cph-dh136)
[Report Description](https://github.com/datsoftlyngby/soft2017fall-lsd-teaching-material/blob/master/assignments/08-Project_report.md)

## Abstract
How to development large systems. We set out to find out. We underwent a project spanning approximatly 9 weeks of development and 6 weeks of maintenence. Here in, topics such as: Scaling, monitoring, development colaboration, security, Continius delivery / Dev ops, logging, deployment strategies and more is touched to accomplice the goal of developing and maintaining large systems. The project was based on a description, which was to make a clone of an existing service, Hackernews. We were tasked with requirements suchs as 95% uptime and not to loose any content from a simulator and more. We managed to do this task, with satisfied results according to self evaluation. 

## Table of Content
- [Requirements, architecture, design and process](#requirements-architecture-design-and-process)
>- [System requirements](#system-requirements)
>- [Development process](#development-process)
>- [Software architecture](#software-architecture)
>- [Software design](#software-design)
>- [Software implementation](#software-implementaion)
- [Maintenance and SLA status](#maintenance-and-sla-status)
>- [Hand-over](#hand-over)
>- [Service-level-agreement](#service-level-agreement)
>- [Maintenance and reliability](#maintenance-and-reliability)
- [Discussion](#discussion)
>- [Technical Discussion](#technical-discussion)
>- [Group work reflection & Lessons learned](#group-work-reflection--lessons-learned)

## Requirements, architecture, design and process
Kristian
```
This section describes the requirements of the project, the architecture that you have chosen,  
the design of the actual components and the process you undertook to build them in the end.
```
#### System requirements
Kristian
```
This section should contain an elaborate description of the requirements for the project.  
This includes the scope of the Hackernews clone (what should it be able to do / what should it not be able to do).
```
The system was a minimal functional clone of the original Hackernews website, which was a system that allowed users to share and discuss stories with a focus on programming and information systems, the system allows self regulation by allowing users to increase the visibility of some discussions, and for long time users to decrease the visibilty of others.  
The system had to handle multiple users posting stories and comments at the same time, while also having a minimum 95% uptime even while part of the system was down for upgrading. The system had to allow users to make a program that can simulate user interaction that creates stories and comments using a REST API, also to query the latest ingested story. likewise the users should also be able to do these actions using a web browser as well.

#### Development process
Kristian - We delegated responsibility. The roles

```
In this part you should show off by telling us all you know about software development processes  
and describe which concepts you used to structure your development.
```
Our choice for structuring our development process was greatly affected by our team size, as a 3 member team we had to pull more than the average 4-5 man team we were suppose to be in, but we took the challenge to better ourself as we had to be more invovled with all parts of the system. We choose to run a mutated version of Scrum as the short development time and iterative development process was needed for a project with this short development cycle.  
we could have gone with the waterfall model, but since our lectures would give new information on a weekly basis that we had to implement into the project this model wouldn't allow for such implementation and was discarded, likewise the Unified process also discarded but not for the same reason, the unified process does allow for iterative development, but at the time we started the project we weren't comepletly familliar with that development process, and we wouldn't get the lecture until 3 days after we passed on the project for remote testing by another group, this made us choose scrum as it allowed us to keep ourself up to date on a day to day basis, and realocate resources depending on the ever changing enviorment, consisting of different courses and lecture material.

#### Software architecture
Kristian
```
In this section you illustrate and describe the architecture of your Hackernews clone.  
That is, you describe how your system is structured and how the different  
parts interact and communicate with each other.
```
Theres two different version of our architecture, doing the development we where promted to extend the system into multiple sub systems, as a practice in low cupled development, the original design is shown in this picture.
![](https://github.com/DanielHauge/HackerNews-Grp8/blob/master/Documentation/Old%20Diagram.jpg)
The original Idea was to have a REST API service, that was designed to handle the specific commands our teachers simulator program would use, the plan was to design out website around these commands, so we didn't have to make multiple commands that would reduce the overall performance of the system. One requirement was to never lose any messages that was recieved, even when the database goes down, to avoid this we implemented our recently gained knowledge of message brokers from the System Intergration course, and designated a MSMQ service called RabbitMQ to handle the messages, another reason for this design was to reduce the load on the API, as we could design it to simple recieve the message, then send it onwards to the save place of the RabbitMQ service, this allowed us to withdraw any database insertion logic from the API so it could handle the messages faster.  
we then had two seperate services that each had a purpose, thier purpose was to handle the messages in the message broker and insert them as fast as they could into the database, or handle the errors should there be any.  
Finally we had a database remotely from the system, this choice was made because the requirements hinted at database upgrade/downtime, so we expected some sort of trouble on the server that the database ran on, and by having it seperate we had a much clearer overview of it.  

Later in the project we were told to expand the model into a lower coupled system, the requirements was to have atleast 3 different systems, which we had at the time, but this prompted a response in our design that changed the architecture model into the following picture.
![](https://github.com/DanielHauge/HackerNews-Grp8/blob/master/Documentation/Componentdiagram2.png)
many of the systems are the same, but 1 key note is the seperation of the API into 2 different ones, our logic was to reduce the delay even further on the API that the teachers simulator program would be communicating with, so we added a identical API with minor tweaks that handled the user interaction from the website. The 2 subsystems that handled the transaction from RabbitMQ to Database got merged into one singular system, with the idea that we would be able to run as many of this subsystem as needed if the transfer rate was suboptimal.  

Present time:

#### Software design
Daniel
```
Here you should sketch your thoughts on the software design before you started  
implementing the system. This includes describing the technical concerns you had  
about the system before you started development, together with all the technical  
components you came up with to fix these concerns and meet the requirements.
```

#### Software implementation
Emmely
```
This section should describe your actual implementation. Mainly how well you  
followed the requirements, process and software design you began with. If your  
system changed during this phase you should summarise the unexpected  
events/problems and explain how you solved them.
```

## Maintenance and SLA status
Daniel
```
This section describes the process of maintaining the software over time,  
starting from the hand-over to the shutting down of your system. The section  
should be written from the viewpoint of the operator, not the developers.
```

#### Hand-over
Emmelys comment. I can write about this
```
In this part, you should describe the hand-over of the system you were operating.  
Specifically you should comment on the quality of the documentation you received  
and whether you felt well equipped to maintaining the system.
```
We thought that the [documentation](https://github.com/HackerNews-lsd2017/hacker-news/wiki) we got was very adequate, well documented in a simple way, easy to read but still content full.
We felt very equipped to monitor their system from the start.
We quickly headed over to the system to look for issues to report.
Our team reported in total around 20 issues. The issues were related to the system as a whole. Both requirements in the use cases that were missing and reporting when monitoring systems went down.
We also talked to the development team directly saying that we thought some Use Cases weren't implemented and agreed that we would give them some time to implement those.
The development team was very quick most of the time and polite answering us and indicating that they took our reported issues seriously.
The communication with the team has been easy both face to face and via email and social media.
A couple of times the development team notified us in advance if there were some issues we need to know about.

#### Service-level agreement
We made a service level agreement with another group acting as operators. The full SLA can be found [Here](https://github.com/DanielHauge/HackerNews-Grp8/wiki/Service-Level-Agreement) in the github wiki section.

But to summerize a few parts of the SLA, here are some agreements:
Availability:
- Meeting Uptimes for Core API (+95% uptime)
- Meeting Uptimes for all systems (+90% uptime)
- Email support: Monitored 12:00 to 20:00 Tuesday - Thursday
- Emails received outside of available hours will be collected, however no action can be guaranteed until the next working day
Service requests:
- 24 hours (during non workdays: Tuesday-Thursday) for issues classified as High priority.
- Within 48 hours for issues classified as Medium priority.
- Within 5 working days for issues classified as Low priority.

To delve deeper into the SLA, it is advised to read the full version: [Here](https://github.com/DanielHauge/HackerNews-Grp8/wiki/Service-Level-Agreement)

Other than that, there were no disagreement from the first itteration of the SLA, and was signed at first proposal.

#### Maintenance and reliability
Daniel
```
This part should contain a description on how you experienced the actual operation.  
Explain how you actually monitored the system to ensure that the SLA was upheld, and  
describe any incidents you experienced that broke (or could potentially break) the SLA.  
Remember to include documentation for each incident! Finally you should conclude  
how well the developers responded to your issues and conclude on how reliable the  
system was overall.
```

## Discussion
```
...
```

#### Technical discussion
Kristian
```
This part summarises both the first and second part of the report by giving an overview  
of the good and bad parts of the whole semester project. Be critical and honest.
```

#### Group work reflection & Lessons learned
```
Give a short reflection on what were the three most important things you learned  
during the project. The lessons learned are with regards to both, what worked well  
and what worked not well. These reflections can cover anything from the sections  
above. That is, development process, architectural and design decisions,  
implementation, maintenance, etc. If you chose to use roles (project manager,  
architect, devops etc.) you should use those to reflect on whether they  
improved the process or not.

Additionally, focus on both, your work as developers as well as operators.
```
