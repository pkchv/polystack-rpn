import sys
from environs import Env

env = Env()
app_id = env('APP_ID', 'worker')

try:
    nats_uri = env('NATS_URI')
    sub_subject = env('SUB_SUBJECT')
    pub_subject = env('PUB_SUBJECT')
except Exception as e:
    print('Required environment variables are missing', file=sys.stderr)
    print(e, file=sys.stderr)
    exit(-1)
