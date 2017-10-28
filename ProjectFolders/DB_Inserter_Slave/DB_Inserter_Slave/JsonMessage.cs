using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace DB_Inserter_Slave
{
    public class JsonMessage
    {
        public string username;
        public string post_type;
        public string pwd_hash;
        public string post_title;
        public string post_url;
        public int post_parent;
        public int hanesst_id;
        public string post_text;
    }
}