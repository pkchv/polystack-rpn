require 'logger'
require 'eventmachine'
require 'nats/io/client'
require './config'

logger = Logger.new(Config::LOG_PATH, formatter: proc { |severity, datetime, app_id, msg|
    date_format = datetime.strftime("%m-%d %H:%M:%S.%3N")
    timestamp = datetime.strftime("%s")
    "#{Config::APP_ID} #{severity} #{date_format} #{timestamp}: #{msg}\n"
})

nats = NATS::IO::Client.new
nats.connect(:servers => [Config::NATS_URI])
logger.info "Connected to #{nats.connected_server}"

EventMachine.run {

    nats.on_error { |e|
        logger.error "Error: #{e}"
    }
  
    nats.on_reconnect {
        logger.info "Reconnected to server at #{nats.connected_server}"
    }
  
    nats.on_disconnect {
        logger.info "Disconnected!"
    }

    nats.subscribe(Config::REQ_ENDPOINT) { |msg, reply, subject| 
        logger.info "Receiving a message from '#{Config::REQ_ENDPOINT}'"
        logger.info "Message payload: #{msg}"
        logger.info "Publishing a message to '#{Config::REQ_WORKER}'"
        logger.info "Message payload: #{msg}"
        nats.publish(Config::REQ_WORKER, msg)
    }

    nats.subscribe(Config::RES_WORKER) { |msg, reply, subject|
        logger.info "Receiving a message from '#{Config::RES_WORKER}'"
        logger.info "Message payload: #{msg}"
        logger.info "Publishing a message to '#{Config::RES_ENDPOINT}'"
        logger.info "Message payload: #{msg}"
        nats.publish(Config::RES_ENDPOINT, msg)
    }

}
