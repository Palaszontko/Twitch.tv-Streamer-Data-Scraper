
# Twitch.tv Streamer Data Scraper

## Overview
The Twitch.tv Streamer Data Scraper is a tool designed to collect and store data from Twitch.tv streamers. 

## Features
- Scrapes data from Twitch.tv streamers.
- Stores scraped data in a structured format.
- Configurable settings for custom scraping needs.
- Logs activities and errors for monitoring and debugging.

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/Palaszontko/Twitch.tv-Streamer-Data-Scraper.git
   cd Twitch.tv-Streamer-Data-Scraper
   ```
2. Install dependencies:
   ```bash
   go mod download
   ```
3a. Run the project:
   ```bash
   go run cmd/main/main.go
   ```
3b. Build and run the project:
   ```bash
   go build -o cmd/main/main.go scraper
   ./scraper
   ```

## Configuration
Configure the scraper by modifying the `configs/config.json` file:
```json
{
  "timePeriod": 365
}
```
- `timePeriod`: Time period of scraped data.
  
## Logging
Logs are stored in the `logs/` directory. Monitor logs for information on scraping activities and errors.

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request with your changes.

## Disclaimer
This script is for educational purposes only. The use of automated tools to create accounts on websites may violate the terms of service of those websites. Use at your own risk.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.


