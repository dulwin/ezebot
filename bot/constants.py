from slacker import Slacker
from pyslack import SlackClient

BOT_ID = 'U1E7ARKRA'
BOT_NAME = 'jarvis'

SLACK_TOKEN = 'get-from-slack'
SLACK = Slacker(SLACK_TOKEN)
PY_SLACK = SlackClient(SLACK_TOKEN)
