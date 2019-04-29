FROM python:3.8.0a3-alpine3.9

# Set in-container working directory
WORKDIR /usr/src/polystack-worker

# Copy dependency definition file
COPY requirements.txt ./

# Install dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Copy project files
COPY . .

# Run the app
CMD ["python", "./main.py"]
