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

#### System requirements

The system was a minimal functional clone of the original Hackernews website, which was a system that allowed users to share and discuss stories with a focus on programming and information systems, the system allows self regulation by allowing users to increase the visibility of some discussions, and for long time users to decrease the visibilty of others.  
The system had to handle multiple users posting stories and comments at the same time, while also having a minimum 95% uptime even while part of the system was down for upgrading. The system had to allow users to make a program that can simulate user interaction that creates stories and comments using a REST API, also to query the latest ingested story. likewise the users should also be able to do these actions using a web browser as well.

#### Development process

Our choice for structuring our development process was greatly affected by our team size, as a 3 member team we had to pull more than the average 4-5 man team we were suppose to be in, but we took the challenge to better ourself as we had to be more invovled with all parts of the system. We choose to run a mutated version of Scrum as the short development time and iterative development process was needed for a project with this short development cycle.  
we could have gone with the waterfall model, but since our lectures would give new information on a weekly basis that we had to implement into the project this model wouldn't allow for such implementation and was discarded, likewise the Unified process also discarded but not for the same reason, the unified process does allow for iterative development, but at the time we started the project we weren't comepletly familliar with that development process, and we wouldn't get the lecture until 3 days after we passed on the project for remote testing by another group, this made us choose scrum as it allowed us to keep ourself up to date on a day to day basis, and realocate resources depending on the ever changing enviorment, consisting of different courses and lecture material.

#### Software architecture

To give a better overview while reading about the Architecture, we'll introduce the system in a short connected manner, with picture reference to help you reader to better navigate this system.  
The system consist of 1 Core API that handles all activity from our teachers simulation program, it's connected to a RabbitMQ that acts as a message buffer and the database which it only pulls data from. We also have a Angular website, which does all it's backend interaction through another API dedicated for website activity, which also connects to the RabbitMQ and only pull data from the database. Our Database Inserter connects to the RabbitMQ and the Database and acts as our message handler logic and database inserter. The database runs seperate from the other systems on another server and acts as a database.
![](https://github.com/DanielHauge/HackerNews-Grp8/blob/master/Documentation/Componentdiagram2.png)

The architecture is located on a Ubuntu server hosted on the digital ocean service, it's main component is a REST API, which was designed to handle the specific commands our teachers simulator program would use, the REST API takes all messages that it receives and send them over to a RabbitMQ message broker channel, it then returns 200 response to the teachers program to conclude the transaction. this was done to reduce the amount of commands that the REST API handles, which allowed us to improve the overall performance of the REST API. The REST API also connects to the database to get the information our teachers request, as they would request the latest harnest ID.  
One requirement was to never lose any messages that was recieved, even when the database goes down, to avoid this we implemented our recently gained knowledge of message brokers from the System Intergration course, and designated a message broker service called RabbitMQ to handle the messages in channels. Our REST API quickly sends the message into the specified channel, where the RabbitMQ acts as a buffer between getting the message and inserting the message, messages will never go into the database if it's down, but we can still receive messages, we wont loose the messages in the buffer if the database is down and the REST API goes down aswell, since the messages are in the RabbitMQ.  
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

Our system consists of the following software:

###### Front end software

**Web application** - written in TypeScript. The framework is Angular2+. 

[Documentation](https://github.com/DanielHauge/HackerNews-Grp8/wiki/Angular-Web-Application)

[Source code](https://github.com/DanielHauge/HackerNews-Grp8/tree/master/Angular)

###### Backend


**Core REST API** - written in Go

[Documentation](https://github.com/DanielHauge/HackerNews-Grp8/wiki/Core-API)
[Documentation](https://github.com/DanielHauge/HackerNews-Grp8/wiki/Go-API%27s)

[Source code](https://github.com/DanielHauge/HackerNews-Grp8/tree/master/ProjectFolders/GoAPI%20-%20Core%20-%20Simulator)


**Web REST API**  - written in Go

[Documentation](https://github.com/DanielHauge/HackerNews-Grp8/wiki/Web-API)

[Source code](https://github.com/DanielHauge/HackerNews-Grp8/tree/master/ProjectFolders/GoApi%20-%20Website) 


**DBInsert slave** - written in Go

[Documentation](https://github.com/DanielHauge/HackerNews-Grp8/wiki/DB-Inserter-Slave)

[Source code](https://github.com/DanielHauge/HackerNews-Grp8/tree/master/ProjectFolders/DB_Inserter_Slave)


**DBInsert slave for Invalid Messages Channel** - written in Java

[Documentation]()

[Source code](https://github.com/DanielHauge/HackerNews-Grp8/tree/master/ProjectFolders/Java-ErrorInserter)


##### How well did we follow the requirements we began with?

For the feedback we got from Helge from our first hand-in of the Requirements and Analysis Document, RAD, we should not invent requirements and this was something we took to our heart.
For this reason, we have put only the minimum and the requirements were fully implemented from the hand-in on November 2.

##### How well did we followed the process we began with


We agreed upon having a project management system. For the team where we would set up the task that needs to be done.
ZenHub was a free online tool that we set up. It reminded of Atlassian JIRA. A project management one of the group members had good experience of.
Unfortunately, we found out that the system had some bugs so we could not use.
We went over to Trello. In this tool, we set up the functionality and overall tasked need to be implemented to implement the system.
This was also where we delegated tasked for in areas we felt confident.
A division of backend and frontend responsibility was because we had different areas we felt we were good in. 
After a few weeks with Trello, we felt it didn't give us enough overview of which tasks were worked on and whos doing what.
We didn't set up a notification system so we all went to the project management system manually.
During the whole project and semester, our group has been using Discord.
We went over to using only Discord. We were three group members that use good communication as a project management tool.
We even use Discord to pin important messages, Voice chat, saving credentials and direct communication (during all hours of the day).

In retrospect, we can attribute much of our project success to a good communication tool and our dedication to it.
The downside by dividing up task according to areas we are comfortable with didn't force us outside our comfort zone.

[Emmely TODO Make statistics of Discord]

##### How well did we followed the software design we began with

Or design was to build a REST API, Messaging Queue, Database and Web Application.
During what we could call elaboration phase we decided to divide up the REST API into two task. 
We incorporated this idea into the design. One handling the Web Application and one handling the Simulator program.
In retrospect, this was a good idea. For the simulator, we wanted to get a story but for the Web Application request, we wanted more information around those stories such as the number of comments, karma and time in "minutes ago" etc.
Those calculations were a lot heavier. So in the Web API we did the heavier calculation to serve a more user-friendly website.
When we handed in, hand in the system to the stakeholder fulfilling all the requirements. 

##### Unexpected events/problems & Solution to the unexpected events/problems (a summarise)


##### System changed during the maintenance phase

###### Front end software

**Web application** - written in Angular2+ . 
Very few changes were made after the hand-in and in November 2.
- A test with implementing front end monitoring/logging with Rollbar.

###### Backend

**Web REST API and Core REST API**  - written in Go
- We added functionality exposing metrics endpoints for Prometheus.
- Added our own way of [logging](https://github.com/DanielHauge/HackerNews-Grp8/blob/master/ProjectFolders/GoApi%20-%20Website/Logger.go).

**DBInsert slave** - written in Go

The DBInserter software was written initially in C#. 
The DBInserter is responsible for handling the messages(request) in the queue and inserting them into the database. 
Because we were using a Relational Database (without foreign keys) the DB inserter quickly (after just a few days with the simulator running) became our bottleneck.
We made a few tweaks to the software to update it. 
But after spending too much time on this issue we replaced it with a new software written in Go that handles the insertions differently.
The Go program hold the Database in memory so that insertions can be done directly without joins.

**DBInsert slave for Invalid Messages Channel** - written in Java

This program came to live after a few weeks into the project.
We started to notice that the simulator send messages that our system couldn't handle and that this often lead to a crash of our system.
For this, we needed to create a new Inserter that could insert those invalid messages in a separate table in the database. We wanted to persist them even though they were invalid.



## Maintenance and SLA status

#### Hand-over

Group I gave us a link to their GitHub repository with a Wiki page, e-mailaddress and where to look for Grafana, their monitoring of the system. 
We thought that the [documentation](https://github.com/HackerNews-lsd2017/hacker-news/wiki) we got was very adequate, well documented in a simple way, easy to read but still content full.
We felt very equipped to monitor their system from the start.

#### Service-level agreement
We made a service level agreement with another group (Group I) acting as operators. The full SLA can be found [Here](https://github.com/HackerNews-lsd2017/hacker-news/wiki/SLA) in the github wiki section.

But to summerize a few parts of the SLA, here are some agreements:

Availability:
- Minimum expected 95% per month.

Responsetime:
- maximum responsetime of service is 60 milliseconds.

To delve deeper into the SLA, it is advised to read the full version: [Here](https://github.com/HackerNews-lsd2017/hacker-news/wiki/SLA)

Other than that, there were no disagreement from the first itteration of the SLA, and was signed at first proposal.

#### Maintenance and reliability

Both us as operators and the developers of the system, was initially confused about how much access we were suppose to have access to, but the developers was quick both in written format and personal meetings, allowing for communication with us operators, this allowed for quick and responsive reactions to every incident that happened, doing the time we as operators where tasked to monitor the system.  
We implemented alerts for when the system would go down, but our primary way of operation was to actively passive monitor, as in we would on regular occasions take a look on the passive monitor system Grafana, allowing us to catch abnormalities that would happen without our or the developers knowledge.  
Although some incidents havn't been reported in written format, as we in many cases reported breaches of the Service level agreement in person, we will show the following cases we did write down and explain each case individually.  

to save space in the report you can read about all 19 written issues [here](https://github.com/DanielHauge/HackerNews-Grp8/blob/master/docs/Issue%20appendix.md)

Concluding on the developers cooperation.  
The group we were operators for where actively discussing and reacting in a very consistent and fast manner, their reaction time and positive feedback to our reports both written or verbal where fast, and their grafana system proclaims 94.964% however it doesn't showcases the general uptime from the lost data in the earlier stages, which was above 95% active state.  

## Discussion

#### Technical discussion

###### First part: the good
Angular. < need help here  
Doing our initial design phase we were contemplating upon how much we would focus, since we did feel confident in how much we could achieve due to our combined enthusiasm, but when we we presented our early version of our RAD (requirements analyses document) to our teacher for feedback, we were told to tone down our expectation, which lead to a minimalistic approach to the requirements, this was absolutely a good direction as we were only 3 in the group where it was expected we were 4-5.  
From the very beginning we used a communication application called discord, the reasoning was for higher level connection even when we weren't at school, the application had both text and voice communication and allowed for us to take notes and share these. This was crucial to our success as a 3 man group, since the downtime between rapid changes and our mutated scrum model, allowed for rapid response to changes and continues discussion and development from the comfort of our homes.  
Our architecture was made very solid from the start due to our lectures in System Intergration, the lessons in message brokers and low coupeling development, allowed for a very flexible development cycle, as we could and did change individual components doing the LSD projects lifespan, another benefit with low coupling, was our ability to implement multiple lanquagde from each members preference, this allowed our 3 man group to work at optimal levels doing the project.  
Software tools like Git and jenkins also helped immensely, as they allowed us to structure and deploy changes, without having a single person dedicate too much time, which could have resulted in our workforce getting cut down to 2 members otherwise.  

###### First part: the bad
We promised to have everyone in the group touch every part of the project, so that we all could improve our skills in all subject matters, since we were 3 member team, each member needed to know how the project worked, and we wanted everyone to specifically be able to recreate the system single handedly. Time restraint trashed that dream, we were unable to delegate the varied task doing our development of the project, as the time contraints and the 3 man size couldn't keep up with the expected 4-5 man development performance.  
Our limited experience with databases and specifically relational database, was a huge problem for us doing the development, we where overestimated how much trouble the database would give us, and we didn't know how to implement it correctly, lacking forign keys and relations in a relational database, resulted in extreme expenses on hosting the database on a digital ocean service, while also proving quite a challenge in making temporary solutions to persistent problems, that we struggled with troughout the entire projects lifespawn.  

###### Second part:the good
Our initial introduction to the group we would be monitoring was fantastic, they were both friendly and intend on a high level communication level throughout the projects lifespawn.  
Their monitoring systems setup had a great overview, which gave a excellent view of their systems performance, allowing us the needed tools to monitor and alert them when needed. Thier response rate to all issues we presented both oral and written was fast, and their communication on each issues was informative, which helped the positive and constructive mood between them as developers and us as operators.  

###### Second part: the bad
When we recieved thier project and the system level agreement, their project wasn't excatly finished, and they did break their system level agreement a few times due to these issues.  

#### Group work reflection & Lessons learned
Maintenance is crucial for large systems, both for developers peace of mind and productivity, and also to maintain a contract between developer and customer. Quality of the product is highly affected by it's performance, and generally the most cost of a product comes from maintaining it after launch, this can be reduced by implementing a strong foundation for monitoring the systems performance.  
System intergration is very powerful for product flexibility, as it allows ongoing improvement and replacement of components after launch, this has great cost initially for it's implementation, but pays back immensly in the maintenance stage of a product.  
Reckless implementation or use of tools and features can cripple products, although time constraints might force these scenarios, it's important to focus on studying the tools, so that you can either replace or improve them before they do damage to the product. 
