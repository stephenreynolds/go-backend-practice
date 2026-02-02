Console.Write("What is your name? ");
string? name = Console.ReadLine();

if (name is null)
{
    Console.WriteLine("Nothing was read. Exiting.");
    return;
}

Console.WriteLine($"Hello, {name}!");

var datetime = DateTime.Now;
Console.WriteLine(
    $"Today is {datetime.ToString("dddd, MMMM d, yyyy")} at {datetime.ToString("h:mm tt")}."
);
