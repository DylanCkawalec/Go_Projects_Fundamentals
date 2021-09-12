import Constants as keys

from telegram import *
from Telegram.ext import*
import Responces as R

print("Bridgy is here")

def start_command(update, context):
    update.message.reply_text('write something and I will do my best to help you')

def help_command(update, context):
    update.message.reply_text('if you need more help, Google is a good friend of mine, ask them')

def handle_message(update, contect):
    text = str(update.message.text).lower()
    responce = R.sample_responses(text)

    update.message.reply_text(responce)

def error(update, contect):
    print(f"update {update} caused error {context.error}")

def main():
    updater = Updater(keys.API_KEY, use_context = True)
    dp = updater.dispatcher

    dp.add_handler(CommandHandler("start", start_command))
    dp.add_handler(CommandHandler("start", help_command))
    dp.add_handler(MessageHandler(Filters.text, handle_message))
    dp.add_error_handler(error)

    updater.start_polling(60)
    updater.idle()

main()