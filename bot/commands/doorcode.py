from constants import PY_SLACK, BOT_NAME

class Doorcode():
    def __init__(self, s):
        self.s = s

    def post_message(self, m):
        PY_SLACK.chat_post_message(
            m.get('channel'),
            'The doorcode is: `{0}`!'.format(123456),
            username=BOT_NAME
        )