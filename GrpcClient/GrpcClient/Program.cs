using System;
using System.Threading.Tasks;
using Grpc.Core;
using Grpc.Net.Client;
using csharpGrpc;

namespace GrpClient
{
    class Program
    {
        static async Task Main(string[] args)
        {
            AppContext.SetSwitch(
                "System.Net.Http.SocketsHttpHandler.Http2UnencryptedSupport", 
                true
            );

            var channel = GrpcChannel.ForAddress("http://localhost:50051", new GrpcChannelOptions
            {
                Credentials = ChannelCredentials.Insecure
            });

            var client = new Users.UsersClient(channel);

            while (true)
            {
                Console.WriteLine("\nEnter an integer between 1-3");
                Console.WriteLine("1 : Search information");
                Console.WriteLine("2 : Add new User");
                Console.WriteLine("3 : Update User");
                Console.WriteLine("4 : Delete User");
                Console.WriteLine("5 : Exit\n");

                int x = Convert.ToInt32(Console.ReadLine());
               
                if (x == 1)
                {
                    Console.WriteLine("Enter User Id\n");
                    Int64 id = Convert.ToInt64(Console.ReadLine());
                    var input = new Request { Id = id };
                    var reply = await client.searchInfoAsync(input);

                    if (!reply.Valid)
                    {
                        Console.WriteLine("No user Found\n");
                        continue;
                    }

                    Console.WriteLine(reply);
                }
                else if (x == 2)
                {
                    Console.WriteLine("Enter \nFirst Name\nLast Name\nAge\nAddress\n");
                    string firstName = Console.ReadLine();
                    string lastName = Console.ReadLine();
                    Int64 age = Convert.ToInt64(Console.ReadLine());
                    string dateJoined = DateTime.Now.ToString();
                    string address = Console.ReadLine();

                    var input = new newUserRequest 
                    {   
                        FirstName = firstName, 
                        LastName = lastName, 
                        Age = age, 
                        DateJoined = dateJoined, 
                        BillingAddress = address 
                    };

                    var reply = await client.AddUserAsync(input);
                    Console.WriteLine(reply);
                }
                else if(x == 3)
                {
                    Console.WriteLine("Enter ID of user to be updated\n");
                    Int64 id = Convert.ToInt64(Console.ReadLine());
                    Console.WriteLine("Enter\nFirst Name\nLast Name\nAge\nAddress\n");
                    string firstName = Console.ReadLine();
                    string lastName = Console.ReadLine();
                    Int64 age = Convert.ToInt64(Console.ReadLine());
                    string address = Console.ReadLine();

                    var input = new updateUserRequest
                    {
                        Id = id,
                        FirstName = firstName,
                        LastName = lastName,
                        Age = age,
                        BillingAddress = address
                    };

                    var reply = await client.UpdateUserAsync(input);
                    if (!reply.Valid)
                    {
                        Console.WriteLine("No user found with the given ID\n");
                        continue;
                    }

                    Console.WriteLine(reply);
                }
                else if(x == 4)
                {
                    Console.WriteLine("Enter ID\n");
                    Int64 id = Convert.ToInt64(Console.ReadLine());

                    var input = new deleteUserRequest
                    {
                        Id = id
                    };

                    var reply = await client.DeleteUserAsync(input);
                    Console.WriteLine(reply);
                }
                else if(x == 5)
                {
                    Console.WriteLine("Closing...\n");
                    break;
                }
                else
                {
                    Console.WriteLine("Enter Valid Integer\n");
                }
            }
        }
    }
}
