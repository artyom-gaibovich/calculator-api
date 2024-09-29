# calculator-api


### Local Setup Instructions for `calculator-api`

Follow these steps to set up and run the `calculator-api` project locally.

#### Prerequisites
1. **Go** installed (version 1.18 or higher).
2. **Docker** and **Docker Compose** installed.

#### Steps

1. **Clone the Repository**
   Make sure to clone the repository to your local machine:
   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. **Set Up Environment Variables**
    - The project uses an `.env` file to manage environment variables. If you don’t have an `.env` file, you can create it by copying the sample:
      ```bash
      make init
      ```
    - This will copy the `.env.sample` file into the `build/.env` file for local use.

3. **Run the Application**
    - Before running the application, load the environment variables from `.env` and start the Go application:
      ```bash
      make run
      ```
    - This command will set all the environment variables and run the `calculator-api` service.

4. **Build the Docker Image**
    - To build the Docker image using Docker Compose:
      ```bash
      make build-image
      ```
    - Alternatively, you can build the Docker image manually with a version argument:
      ```bash
      make build-image2
      ```
    - This will use the build argument `VERSION=1.2.3` to tag the Docker image with the name defined in the `Makefile`.

Now your `calculator-api` should be running locally, and you can interact with it via the exposed endpoints.


### API
#### 1.  Health-check handler
###### Request
```bash
GET http://localhost{port}:/health-check
```
###### Response

```json
{"status":"OK","message":"Service successfully worked!"}
```