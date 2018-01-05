3th November Antivirus issue  
![](https://i.gyazo.com/941c6c9ad87b7dbc48c50a5f0f5dc28c.png)  
the first major issue wasn't part of the system level agreement, but the way they setup thier web interface for user interaction of their system had a issue, the issue was they redirected the web client to thier backend to get their real user interface, this is generally known as a malicious action, a website could redirect the user to a http call that would download virus or malware, it took a long time to fix this issue, as it was a direct result of their initial architecture so it required extensive remodelling to solve this case.  

8th November users without password  
![](https://i.gyazo.com/bf869551b151e64cc5fc15dc8f583831.png)  
this issue we spottet was a user creation error, the problem was the user could register without a password, login without a password and do all actions as a verified user, this is generally a issue of verification both on creation and login, since if a user doesn't need password it might be possible to login on another user as the password might be irrelevant.  

8th November no feedback on user creation  
![](https://i.gyazo.com/d11d36f4d274e39077424ba2aa1570bc.png)  
this issue was a minor feedback problem, where you wouldn't know if your request for account creation was succesful or not.  

8th November 404 refresh error  
![](https://i.gyazo.com/a98092f6ceee0da0a6452b8a5d003f41.png)  
Doing a refresh when attempting to login would result in a 404, this problem existed because of their initial architecture, which didn't redirect to sub domains doing each action in the webpage navigation.  

10th November Grafana information confusion  
![](https://i.gyazo.com/fa29b6b5aaa1e53a95c0b4807d3ce97e.png)  
One information graph on their grafana confused us as it showed alot more "request" than we observed the teachers had sent, turns out this graph showed all actions of all systems, resulting in the higher than expected request counter.  

11th November database update  
![](https://i.gyazo.com/b610f0333b9c2982af7f7feebba92e84.png)  
this was a error on our side as operators, as we observed actively at the time and presumed the abnormal activity to be a crash, turned out to be a database upgrade, we would have known it was this if we had checked http://46.101.28.25:8080/status  

12th November website not showing stories  
![](https://i.gyazo.com/6e1603e01b80ffee6020f89dfe09a8cb.png)  
This issue spawned from the previous days database upgrade.  

12th November Service crash  
![](https://i.gyazo.com/664e9a43d11e87a32ff4ab26cc1a6e01.png)  
Further issues from the 11th November update resulted in major changes to their system, which made us notice many crashes and abnormalities.  

13th November Prometheus crash  
![](https://i.gyazo.com/00c37644baade9d88f9f12c22de14bee.png)  
unrelated to the previous days development, this issue happened multiple times doing the semester where we monitored their system, we now know afterwards these crashes where related to running prometheus on a docker container, which they later implemented a volume for.  

13th November Intentional stress test  
![](https://i.gyazo.com/31bde7c8b7b372aca9f5003204a19457.png)  
this issue was part of a assignment for the developers to attempt to break their system and for us as operators to observe and report it, the issue came from them using a simulation that posted as fast as it could to their service, resulting in a abnormal request spike.  

14th November missing features comments 
![](https://i.gyazo.com/8c796c22d5dc0f33a4a4a9ebeea1cddb.png)  
this issue was at the time of reporting it not implemented, this was fixed later down the line.  

14th November missing feature logout  
![](https://i.gyazo.com/bd984487ef196ac8ea760ba69a28fbb2.png)  
this issue was at the time of reporting it not implemented, this was fixed later down the line.  

14th November missing feature update user  
![](https://i.gyazo.com/0e483304c6f9c9ba3335bf7c397a0f43.png)  
this issue was at the time of reporting it not implemented, this was fixed later down the line.  

15th November temporary downtime  
![](https://i.gyazo.com/1d6386ba66942b3f0efe8a5abde1b9b7.png)  
this was a minor hiccup on the serverside, documentation seems to indicate it was temporary and no developer actions where needed.  

16th November Invalid stories  
![](https://i.gyazo.com/e07fa55a9ce08a34d27d94eed613f29a.png)  
Due to a error in implementation on the developers side, the system would display the stories that had it's content deleted, resulting in showcases of invalid stories.  

16th November Grafana crash  
![](https://i.gyazo.com/6ae549bf31b4a930353bff64bed75535.png)  
one of the reported grafana crashes.  

24th November Grafana crash  
![](https://i.gyazo.com/cfb20c983c2f513e87bfc90d1e7c4e39.png)  
This crash wasn't found out by our alerts, at the time of writing this report it seems we forgot to fix our grafana alerts after the previous developments, as it's possible we didn't have alerts from this point onwards.  

29th November OWTF testing  
![](https://i.gyazo.com/d4ee34b6b0c9cc81afa2389c0ec37e3b.png)  
as assignment we had to test our designated system we operated on, so to alert the developers what time we would be doing the testing, we discussed aproximately when we do this and report the beginning and end of the testing, as to allow them to respond in time incase of system failure from the OWTF stress testing.  

11th december Grafana crash  
![](https://i.gyazo.com/48559583616ee590d6b65e6e99e226f8.png)  
last reported issue although not the last grafana crash.
