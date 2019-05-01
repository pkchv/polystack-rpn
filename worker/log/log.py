import logging
from environs import Env

env = Env()
app_id = env('APP_ID', 'worker')
filename = env('LOGGER_FILENAME', './logs/worker.log')
log_format = env('LOGGER_LOG_FORMAT', '%(asctime)s %(name)-12s %(levelname)-8s %(message)s')
date_format = env('LOGGER_DATE_FORMAT', '%m-%d %H:%M')
level = env('LOGGER_LEVEL', 'DEBUG')
logger_filemode = env('LOGGER_LOG_FILE_MODE', 'w')

logging.basicConfig(level=logging.getLevelName(level),
                    format=log_format,
                    datefmt=date_format,
                    filename=filename,
                    filemode=logger_filemode)

log = logging.getLogger(app_id)
