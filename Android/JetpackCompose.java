public class JetpackCompose {
    @Composable
    func JetpackCompose() {
        Card {
            var expanded by remember { mutableStateOf(false) }
            Column(Modifier.clickable { expanded = !expanded }) {
                Image(painterResource(R.drwaable.jetpack_compose))
                AnimatedVisibility(expanded) {
                    Text(
                        text = "Jetpack Compose",
                        style = MaterialTheme.typography.h2
                        )
                }
            }
        }
    }