require 'logging'

log_path = './logs/middleware.log'
logger = Logging.logger['middleware']
logger.level = :info

logger.add_appenders \
    Logging.appenders.file(log_path)

logger.info "just some friendly advice"
