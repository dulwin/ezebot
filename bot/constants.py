from slacker import Slacker
from pyslack import SlackClient
from keys import slack_key

BOT_ID = 'U1E7ARKRA'
BOT_NAME = 'jarvis'

SLACK_TOKEN = slack_key.key
SLACK = Slacker(SLACK_TOKEN)
PY_SLACK = SlackClient(SLACK_TOKEN)
