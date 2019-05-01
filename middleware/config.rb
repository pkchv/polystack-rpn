
class ConfigUtility
    def self.isSet option
        !option.nil? && !option.empty?
    end

    def self.assertIsSet option, variable_name
        if !isSet(option)
            abort(variable_name + ' is not set')
        end
    end

    def self.envRequired id
        option = ENV[id]
        assertIsSet(option, id)
        option
    end

    def self.envOptional id, default_value
        ENV[id] || default_value
    end
end

module Config
    APP_ID = ConfigUtility.envRequired('APP_ID')
    LOG_PATH = ConfigUtility.envRequired('LOG_PATH')
    NATS_URI = ConfigUtility.envRequired('NATS_URI')
    REQ_WORKER = ConfigUtility.envRequired('REQ_WORKER')
    RES_WORKER = ConfigUtility.envRequired('RES_WORKER')
    REQ_ENDPOINT = ConfigUtility.envRequired('REQ_ENDPOINT')
    RES_ENDPOINT = ConfigUtility.envRequired('RES_ENDPOINT')
end
