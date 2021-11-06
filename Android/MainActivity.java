class MainActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(MessageCard("Android"))
    }
}

data class Message(val author: String, var body: String)

@Composable
fun MessageCard(msg: Message){
    Column{
        Text(text = msg.author)
        Text(text = msg.body)
    }
}


@Preview
@Composable
fun PreviewMessageCard(){
    // ComposeTutorialTheme {
    //     MessageCard("Android")
    // }
    MessageCard(
        msg = Messag("Colleague", "test")
    )
}