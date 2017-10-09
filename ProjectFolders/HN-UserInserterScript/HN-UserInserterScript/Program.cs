using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using MySql.Data;
using MySql.Data.MySqlClient;
using System.Data;

namespace HN_UserInserterScript
{
    class Program
    {
        static void Main(string[] args)
        {
            IEnumerable<string[]> lines = LoadCsvData(@"C:\Users\Animc\Desktop\users.csv", ',');
            Console.WriteLine("Reader Started");



            using (MySqlConnection connection = new MySqlConnection("server=46.101.103.163;user id=myuser;database=HackerNewsDB;persistsecurityinfo=True;allowuservariables=True;Pwd=HackerNews8"))
            {
                MySqlCommand cmd = new MySqlCommand(
                    "INSERT INTO User (Name, Password, KarmaPoints) VALUES (@Name, @Password, @KarmaPoints)");

                cmd.CommandType = CommandType.Text;
                cmd.Connection = connection;

                Console.WriteLine("SQLCOMMAND PROCEDURE STATED!");
                connection.Open();
                Console.WriteLine("CONNECTION OPENED!");

                if (lines != null)
                {
                    foreach (var line in lines)
                    {
                        cmd.Parameters.Clear();
                        cmd.Parameters.AddWithValue("@KarmaPoints", 0);
                        cmd.Parameters.AddWithValue("@Name", line[0]);
                        cmd.Parameters.AddWithValue("@Password", line[1]);
                        cmd.ExecuteNonQuery();
                        Console.WriteLine("Inserted - User: " + line[0] + " ||| Password: " + line[1]);

                    }
                }
                else { Console.WriteLine("Lines were empty"); }
                connection.Close();
                Console.WriteLine("CLOSING");
            }

        }

        private static IEnumerable<string[]> LoadCsvData(string path, params char[] separator)
        {
            return from line in File.ReadLines(path)
                   let parts = (from p in line.Split(separator, StringSplitOptions.RemoveEmptyEntries) select p)
                   select parts.ToArray();
        }
    }
    
}
