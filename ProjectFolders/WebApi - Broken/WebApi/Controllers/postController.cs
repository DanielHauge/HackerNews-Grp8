using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Net.Http;
using System.Web.Http;
using WebApi.Models;
using System.Diagnostics;

namespace WebApi.Controllers
{
    public class postController : ApiController
    {
        // GET api/<controller>
        public IEnumerable<string> Get()
        {
            return new string[] { "value1", "value2" };
        }

        // GET api/<controller>/5
        public string Get(int id)
        {
            return new SmartDBConnectionClass("server=46.101.103.163;user id=myuser;database=HackerNewsDB;SslMode=none;persistsecurityinfo=True;allowuservariables=True;Pwd=HackerNews8").GetLatestHanesst_ID().ToString();
        }

        // POST api/<controller>
        public void Post([FromBody]PostClass value)
        {
            Debug.Print("UserName: "+value.username + " - Pwd: "+value.pwd_hash);
            //PostClass PostRequest = value;
            SmartDBConnectionClass DB = new SmartDBConnectionClass("server=46.101.103.163;user id=myuser;database=HackerNewsDB;SslMode=none;persistsecurityinfo=True;allowuservariables=True;Pwd=HackerNews8");

            //if (DB.VerifyUser(PostRequest.username, PostRequest.pwd_hash))
            //{  
            //    DB.SavePost(PostRequest);
            //    Debug.Print("Did Verify");
            //}
            //else Debug.Print("did not verify");

            DB.SavePost(value);

        }

        // PUT api/<controller>/5
        public void Put(int id, [FromBody]string value)
        {
        }

        // DELETE api/<controller>/5
        public void Delete(int id)
        {
        }
    }
}