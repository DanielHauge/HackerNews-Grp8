using MySql.Data.MySqlClient;
using System;
using System.Collections.Generic;
using System.Data;
using System.Data.SqlClient;
using System.Linq;
using System.Web;

namespace WebApi.Models
{
    public class SmartDBConnectionClass
    {
        private string constring { get; set; }




        public SmartDBConnectionClass(string connectionString)
        {
            constring = connectionString;



        }

        public bool VerifyUser(string un, string pwd)
        {
            bool res = false;
            MySqlConnection con = new MySqlConnection(constring);
            con.Open();
            var command = new MySqlCommand("SELECT        Password, Name FROM            `User` WHERE(Name = '"+un+"')", con);
            MySqlDataReader dr = command.ExecuteReader();

            while (dr.Read())
            {
                string pass = dr[0].ToString();
                if (pass == pwd) { res = true; }
            }
            con.Close();
            con.Dispose();
            return res;
        }

        public void CreateUser(string username, string password)
        {

            using (MySqlConnection connection = new MySqlConnection(constring))
            {
                MySqlCommand cmd = new MySqlCommand(
                    "INSERT INTO User (Name, Password, KarmaPoints) VALUES (@Name, @Password, @KarmaPoints)");

                cmd.CommandType = CommandType.Text;
                cmd.Connection = connection;

                cmd.Parameters.AddWithValue("@Name", username);
                cmd.Parameters.AddWithValue("@Password", password);
                cmd.Parameters.AddWithValue("@KarmaPoints", 0);

                connection.Open();
                cmd.ExecuteNonQuery();
                connection.Close();
            }

        }

        public void SavePost(PostClass PostToSave)
        {

            using (MySqlConnection connection = new MySqlConnection(constring))
            {
                MySqlCommand cmd = new MySqlCommand("INSERT INTO " + GetPostTable(PostToSave.post_type) +  " " + GetParaMeters(PostToSave.post_type) +   " VALUES " + getValues(PostToSave.post_type));
                //INSERT INTO Thread(Name, UserID, Time, Han_ID) VALUES (@ThreadID, @Comment, @UserID, @Time, @Han_ID)
                cmd.CommandType = CommandType.Text;
                cmd.Connection = connection;

                cmd.Parameters.AddWithValue("@ThreadID", PostToSave.post_parrent);
                cmd.Parameters.AddWithValue("@Comment", PostToSave.post_text);

                
                cmd.Parameters.AddWithValue("@UserID", GetUserID(PostToSave.username));
                cmd.Parameters.AddWithValue("@Time", DateTime.Now);
                cmd.Parameters.AddWithValue("@Han_ID", PostToSave.hanesst_id);
                cmd.Parameters.AddWithValue("@CommentKarma", 0);
                
                cmd.Parameters.AddWithValue("@Name", PostToSave.post_title);
                
                

                connection.Open();
                cmd.ExecuteNonQuery();
                connection.Close();
            }

        }

        private int GetUserID(string username)
        {
            int res = 0;
            MySqlConnection con = new MySqlConnection(constring);
            con.Open();
            var command = new MySqlCommand("SELECT ID FROM User WHERE Name LIKE '" + username + "'", con);
            MySqlDataReader dr = command.ExecuteReader();
            while (dr.Read())
            {
                string resultID = dr[0].ToString();
                res = int.Parse(resultID);
            }
            con.Close();
            con.Dispose();
            return res;
        }

        private string GetParaMeters(string post_type)
        {
            if (post_type == "story") { return "(Name, UserID, Time, Han_ID)"; }
            else if (post_type == "comment") { return "(ThreadID, Comment, UserID, CommentKarma, Time, Han_ID)"; }
            else { return "Invalid"; }
        }

        private string GetPostTable(string post_type)
        {
            if (post_type == "story") { return "Thread"; }
            else if (post_type == "comment") { return "Comment"; }
            else { return "invalid"; }
        }

        private string getValues(string post_type)
        {
            if (post_type == "story") { return "(@Name, @UserID, @Time, @Han_ID)"; }
            else if (post_type == "comment") { return "(@ThreadID, @Comment, @UserID, @CommentKarma, @Time, @Han_ID)"; }
            else { return "Invalid"; }

        }

        public int GetLatestHanesst_ID()
        {
            int res = 0;

            MySqlConnection con = new MySqlConnection(constring);
            con.Open();
            var command = new MySqlCommand("SELECT MAX(Han_Id) FROM Comment FULL OUTER JOIN Thread", con);
            MySqlDataReader dr = command.ExecuteReader();
            while (dr.Read())
            {
                res = int.Parse(dr[0].ToString());
            }
            con.Close();
            con.Dispose();
            return res;


        }
    }
}