# Import libraries
import requests
from telegram.ext import Updater, CommandHandler, MessageHandler, Filters
import time

# Replace with your actual API key
API_KEY = "5d4500edca427c597ba5339f"
CHAT_ID = "YOUR_TELEGRAM_CHAT_ID"  # Replace with your Telegram chat ID

# Function to check exchange rate
def check_exchange_rate():
  url = f"https://api.exchangerate-api.com/v1/latest?base=USD&symbols=CNY"
  headers = {"apikey": API_KEY}
  response = requests.get(url, headers=headers)
  data = response.json()
  
  # Check if successful
  if data["success"]:
    rate = data["rates"]["CNY"]
    if rate < 7.19:
      message = f"USD to RMB exchange rate dropped below 7.19! Current rate: {rate}"
      send_telegram_message(message)
  else:
    print("API request failed!")

# Function to send message via Telegram bot
def send_telegram_message(message):
  updater = Updater(token="YOUR_TELEGRAM_BOT_TOKEN")  # Replace with your bot token
  dispatcher = updater.dispatcher
  
  def send_message(update, context):
    context.bot.send_message(chat_id=CHAT_ID, text=message)
  
  dispatcher.add_handler(MessageHandler(Filters.text & ~Filters.command, send_message))
  updater.start_polling()
  updater.idle()

# Run the check function periodically (e.g., every hour)
if __name__ == "__main__":
    while True:
        check_exchange_rate()
        # Adjust the delay as needed (in seconds)
        time.sleep(3600)
