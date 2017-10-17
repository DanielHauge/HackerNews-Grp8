using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;

namespace WebApi.Models
{
    public class PostClass
    {
        public string username { get; set; }
        public string post_type { get; set; }
        public string pwd_hash { get; set; }
        public string post_title { get; set; }
        public int post_parrent { get; set; }
        public int hanesst_id { get; set; }
        public string post_text { get; set; }
    }
}