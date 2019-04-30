FROM ruby:2.6.3

# throw errors if Gemfile has been modified since Gemfile.lock
RUN bundle config --global frozen 1

# Set in-container working directory
WORKDIR /usr/src/app

# Copy dependency definition files
COPY Gemfile Gemfile.lock ./

# Install dependencies
RUN bundle install

# Copy project files
COPY . .

# Run the app
CMD ["ruby", "./main.rb"]
