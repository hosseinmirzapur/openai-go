Sure! Here's an example of a README.md file for your CLI app developed with Golang:

# CLI App for OpenAI Services

This is a command-line interface (CLI) application that utilizes the OpenAI APIs to provide chat completion, image generation, and audio transcription services. It allows you to interact with the OpenAI models programmatically through a simple and convenient CLI interface.

## Features

- **Chat Completion**: Use OpenAI's chat models to generate human-like responses to user queries.
- **Image Generation**: Generate images based on provided prompts using OpenAI's image generation models.
- **Audio Transcription**: Transcribe audio files into text using OpenAI's audio transcription API.

## Requirements

- **go**: should be version 1.20 or higher
- **OpenAI API Key**: can be obtained from your openai account
- **git**: If not installed, you can download zip file of this project

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/hosseinmirzapur/openai-go.git


   ```

2. Navigate to the project directory:

   ```bash
   cd openai-go


   ```

3. Install the required dependencies:

   ```bash
   go mod download


   ```

4. Obtain an API key from OpenAI by signing up on their website (https://openai.com).

5. Set the API key as an environment variable: (or create an .env file and use godotenv)

   ```bash
   export OPENAI_API_KEY=your-api-key

   ```

## Usage

The CLI app can be simply used by running the specific file inside the "cmd" module:

- `chat`: Interact with the chat models. Example usage:

  ```bash
  go run cmd/chat/main.go

  # you will be prompted for a text


  ```

- `generate-image`: Generate an image based on a prompt. Example usage:

  ```bash
  go run cmd/image/main.go

  # you will be prompted for an image description text


  ```

- `transcribe-audio`: Transcribe an audio file into text. Example usage:

  ```bash
  go run cmd/transcribe/main.go

  # you will be prompted for a filepath to your desired audio

  ```

## Configuration

You can customize the behavior of the CLI app by modifying the each service file located at `services/openai_service` directory. The service files allow you to specify default settings for the OpenAI API, such as the model to use, temperature, and max tokens.

## Contributing

Contributions are welcome! If you encounter any issues or have suggestions for improvements, please create a new issue on the project's GitHub repository. Pull requests are also encouraged.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgements

- This CLI app was developed using the Golang programming language.
- It utilizes the OpenAI APIs for chat completion, image generation, and audio transcription.

## Contact

For any inquiries or questions, please contact me via [email](mailto:hosseinmirzapur@gmail.com).

---

Feel free to customize this README.md file according to your specific project requirements and add any additional sections or information that you think would be relevant.
