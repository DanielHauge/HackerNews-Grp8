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
To give a better overview while reading about the Architecture, we'll introduce the system in a short connected manner, with picture reference to help you reader to better navigate this system.  
The system consist of 1 Core API that handles all activity from our teachers simulation program, it's connected to a RabbitMQ that acts as a message buffer and the database which it only pulls data from. We also have a Angular website, which does all it's backend interaction through another API dedicated for website activity, which also connects to the RabbitMQ and only pull data from the database. Our Database Inserter connects to the RabbitMQ and the Database and acts as our message handler logic and database inserter. The database runs seperate from the other systems on another server and acts as a database.
![](https://github.com/DanielHauge/HackerNews-Grp8/blob/master/Documentation/Componentdiagram2.png)

The architecture is located on a Ubuntu server hosted on the digital ocean service, it's main component is a REST API, which was designed to handle the specific commands our teachers simulator program would use, the REST API takes all messages that it receives and send them over to a RabbitMQ message broker channel, it then returns 200 response to the teachers program to conclude the transaction. this was done to reduce the amount of commands that the REST API handles, which allowed us to improve the overall performance of the REST API. The REST API also connects to the database to get the information our teachers request, as they would request the latest harnest ID.  
One requirement was to never lose any messages that was recieved, even when the database goes down, to avoid this we implemented our recently gained knowledge of message brokers from the System Intergration course, and designated a MSMQ service called RabbitMQ to handle the messages. Our REST API quickly sends the message into the specified channel, where the RabbitMQ acts as a buffer between getting the message and inserting the message, messages will never go into the database if it's down, but we can still receive messages, we wont loose the messages in the buffer if the database is down and the REST API goes down aswell, since the messages are in the RabbitMQ.  
We also have a seperate REST API that handles all Web based communication, the choice for seperating the communication of the website and the teachers simulator program was to further spread the load, since we wanted no interference with our teachers simulator and our digestion of the simulation. The web API is connected to a angular website, and has all the features of the other API, so the Angular website uses the same commands, but present a graphical user interface for all web browser users.  
For handling the messages in the RabbitMQ buffer we implemented a database inserter named "DB Inserter slave", it tries to take the first message in the message brokers channel, and process the information so it can insert it into the database, if the message is faulty it sends it to a dead end channel on the message broker, otherwise it'll insert the data into the database as fast as possible.  
Lastly we have the database, which is a mySQL database on a seperate ubuntu server hosted on digital ocean.

#### Software design

We made some basic diagrams to agree upon the systems operations.
- API: [Here](https://github.com/DanielHauge/HackerNews-Grp8/blob/master/Documentation/GoOverview.png)
- Architecture: [Here](https://github.com/DanielHauge/HackerNews-Grp8/blob/master/Documentation/Componentdiagram1.png)
- Data flow: [Here](https://github.com/DanielHauge/HackerNews-Grp8/blob/master/Documentation/Viewthread.png)
- Gui-Prototype: [Here](https://github.com/DanielHauge/HackerNews-Grp8/blob/master/Documentation/Web%20UI%20prototype.png)

Before we started developing, we had some concerns about the 2 different users who needed to interact with the system concurrently. 1 type of user is the human user, they were not restricting us in any way. We can decide what ID,s those users posts, comments etc. has, but for the other type of user, the simulator. The simulator wanted to decide what ID, "Harnest_ID" and "Post_parrent". This would cause a problem, if per say, the website user posted a comment and would occupy that ID, then an error would occur because the simulator insists upon deciding that ID. This was a integration challenge we needed to figure out. How should we fix the issue, that 2 different users where one of them would like to decide identifier beforehand.

The fix to this specific issue was to make 2 different interfaces. We made and decided upon the ID,s that we will be using. We made one interface for the website users, and made one interface (API) that would translate the simulators ID,s into our own ID,s and Thread/Comment-ID,s. This way, we threat the data the same even for 2 different ways of interacting with the system.

#### Software implementation
Emmely
```
This section should describe your actual implementation. Mainly how well you  
followed the requirements, process and software design you began with. If your  
system changed during this phase you should summarise the unexpected  
events/problems and explain how you solved them.
```
###### Front end software

**Web application** - written in Angular2+ . 
[Documentation]()
[Source code]()

###### Backend

**Core REST API** - written in Go
[REST API documentation]()
[Source code]()

**Web REST API**  - written in Go
[REST API documentation]()
[Source code]() 
**DBInsert slave** - written in Go
[Documentation]()
[Source code]()
**DBInsert slave for Invalid Messages Channel** - written in Java
[Documentation]()
[Source code]()

##### How well did we follow the requirements we began with?
Our group but the requirement very low from the start. We got the feedback from Helge from our first hand-in of the Requirements and Analysis Document, RAD, that we should not invent requirements and this was something we took to our heart.
For this reason we have put only the minimum and the requirements were fully implemented from the hand-in on November 2.
##### How well did we followed the process we began with
In the beg
[comment]: <> (I don't know if it's relevant but we used Trello as project management system but went away from that)
[comment]: <> What process?

##### How well did we followed the software design we began with
Or design was to build a REST API, Messaging Queue, Database and Web Application.
During what we could call elaboration phase we decided to divide up the REST API for two task. 
We incorporated this idea into the design. One handling the Web Application and one handling the Simulator program.
In retrospect this was a good idea. For the simulator we wanted to get a story but for the Web Application request we wanted more information around those stories such as number of posts and time etc.
Those calculations were a lot heavier. So in the Web API we did heavier calculation to serve a more user friendly website.
When we handed in, handed in the system to the stake holder fulfilling all the requires. 
We had developed the system as we designed it. With continues delivery 

##### System changed during the maintenance phase
###### Front end software

**Web application** - written in Angular2+ . 
Very few changes were made after the hand-in and in November 2.
- Small bugg with element
- Implementing front end monitoring/logging with rollbar. But not a great tool.
###### Backend

**Core REST API** - written in Go

**Web REST API**  - written in Go
 
**DBInsert slave** - written in Go
The DBInserter software were written initially in C#. 
The DBInserter is responsible for handling the messages(request) in the queue and inserting them into the database. 
Because we were using a Relational Database (without foreign keys) the DBInserter quickly (after just a few days with the simulator running) became our bottleneck.
We made a few tweaks to the software to update it. 
But after spending too much time on this issue we replaced it with a new software written in Go that handles the insertions differently.
By buffering up the insertions into one insert query.
 
**DBInsert slave for Invalid Messages Channel** - written in Java
This program came to live after a few weeks into the project.
We started to notice that the simulator send messages that our system couldn't handle and that this often lead to a crash of our system.
For this we needed to create a new Inserter that could insert those invalid messages. The way it works .......


##### Unexpected events/problems & Solution to the unexpected events/problems (a summarise)

##### Conclusion
The software fontend is not depending on how large the system is. 
If we would have gotten a large number of request to the frontend the way we have would designed it would stay the same. 
The software could not optimise however the server/servers serving the Web Application for hat but our servers. 
Using docker swarm would need to instantiate more servers handling the load. 
There are other factors that determined how to write. 

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
