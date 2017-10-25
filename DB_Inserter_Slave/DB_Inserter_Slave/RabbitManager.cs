using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using System;
using System.Text;
using System.Web.Script.Serialization;
using MySql.Data.MySqlClient;

namespace DB_Inserter_Slave
{
    class RabbitManager
    {
        private string sqlString { get; set; }

        public RabbitManager()
        {
            sqlString = "server = 46.101.103.163; user id = admin; password = hackernews8; database = HackerNewsDB; allowuservariables = True; persistsecurityinfo = True";
        }
        public void InsertMessage(MySqlCommand command)
        {
            MySqlConnection sqlConnection = new MySqlConnection(sqlString);
            sqlConnection.Open();
            command.Connection = sqlConnection;
            command.ExecuteNonQuery();
            sqlConnection.Close();
            sqlConnection.Dispose();
        }

        public void ReceiveMessage(string messageChannel)
        {
            var factory = new ConnectionFactory() { HostName = "138.197.186.82", UserName = "admin", Password = "password" };
            using (var connection = factory.CreateConnection())
            using (var channel = connection.CreateModel())
            {
                channel.QueueDeclare(queue: messageChannel,
                                     durable: true,
                                     exclusive: false,
                                     autoDelete: false,
                                     arguments: null);

                channel.BasicQos(0, 1, false);

                Console.WriteLine(" [*] Waiting for messages.");

                var consumer = new EventingBasicConsumer(channel);
                consumer.Received += (model, ea) =>
                {
                    var body = ea.Body;
                    var jsonmessage = (JsonMessage)new JavaScriptSerializer().Deserialize(Encoding.UTF8.GetString(body), typeof(JsonMessage));
                    //decide where the message go
                    if (jsonmessage.post_type == "story")
                    {
                        MySqlConnection sqlConnection = new MySqlConnection(sqlString);
                        MySqlCommand command = new MySqlCommand("Select ID from HackerNewsDB.User where Name = '" + jsonmessage.username + "'", sqlConnection);
                        sqlConnection.Open();
                        MySqlDataReader reader = command.ExecuteReader();
                        int userID = 0;
                        while (reader.Read())
                        {
                            string result = reader[0].ToString();
                            userID = int.Parse(result);
                        }
                        sqlConnection.Close();
                        sqlConnection.Dispose();
                        Threads thread = new Threads { Name = jsonmessage.post_title, UserID = userID, Post_URL = jsonmessage.post_url, Han_ID = jsonmessage.harnesst_id, Time = DateTime.Now };
                        string message = "Insert into HackerNewsDB.Thread(Name,UserID,Time,Han_ID,Post_URL) values(" + thread.Name + "," + thread.UserID + "," + thread.Time + "," + thread.Han_ID + "," + thread.Post_URL + ")";
                        Console.WriteLine("Thread get");
                        //InsertMessage(""NEWSQLCOMMAND"");
                    }
                    else if (jsonmessage.post_type == "comment")
                    {
                        MySqlConnection sqlConnection = new MySqlConnection(sqlString);
                        MySqlCommand command = new MySqlCommand("Select ID from HackerNewsDB.User where Name = " + jsonmessage.username, sqlConnection);
                        sqlConnection.Open();
                        MySqlDataReader reader = command.ExecuteReader();
                        int userID = 0;
                        while (reader.Read())
                        {
                            string result = reader[0].ToString();
                            userID = int.Parse(result);
                        }
                        sqlConnection.Close();
                        sqlConnection.Dispose();
                        Comment comment = new Comment { UserID = userID, ThreadID = jsonmessage.post_parent, ParentID = jsonmessage.post_parent, Han_ID = jsonmessage.harnesst_id, Time = DateTime.Now };
                        MySqlCommand InsertCommand = new MySqlCommand("Insert into HackerNewsDB.Comment (ThreadID, Name, UserID, CommentKarma, Time, Han_ID, ParentID) values (@parentID, @Name, @UserID, @Number, @Time, @Han_ID, @parentID)");
                        InsertCommand.Parameters.AddWithValue("@parentID", comment.ParentID);
                        InsertCommand.Parameters.AddWithValue("@Name", comment.Name);
                        InsertCommand.Parameters.AddWithValue("@UserID", comment.UserID);
                        InsertCommand.Parameters.AddWithValue("@Number", 0);
                        InsertCommand.Parameters.AddWithValue("@Time", comment.Time);
                        InsertCommand.Parameters.AddWithValue("@Han_ID", comment.Han_ID);
                        Console.WriteLine("Comment get");
                        InsertMessage(InsertCommand);
                    }
                    else if (jsonmessage.post_type == "UserInsert")
                    {
                        //string message = "Insert into HackerNewsDB.User(Name,KarmaPoints,Password,Email) values('" + userMessage.Name + "','" + userMessage.KarmaPoints + "','" + userMessage.Password + "','" + userMessage.Email + "')";
                        //Console.WriteLine("User get");
                        //InsertMessage(message);
                    }
                    else
                    {
                        Console.WriteLine("ERROR");
                    }

                    Console.WriteLine(" [x] Done");

                    channel.BasicAck(deliveryTag: ea.DeliveryTag, multiple: false);
                };
                channel.BasicConsume(queue: messageChannel,
                                     autoAck: false,
                                     consumer: consumer);
                Console.ReadLine();
            }
        }
    }
}
