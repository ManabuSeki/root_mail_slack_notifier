# RootMailSlackNotifier

## DESCRIPTION
Notify slack of mail addressed to Root.

## HOW TO USE

1. Please download the file suitable for your environment.
2. Unzip the downlaod file.
3. Move to this binary file to `/usr/local/bin/slack-notifier`
4. Create file in `/etc/postfix/slack_notice.json`  
    ```json
    {
      "WebhookURL":"Slack_Webhook_URL",
      "Username":"Username",
      "Channel":"general",
      "IconEmoji":":email:",
      "Color":"warning"
    }
    ```
5. postfix aliase setting  
Add the following setting to the bottom line.  
    ```
    # vi /etc/aliases

    root: "| /usr/local/bin/slack-notify -config /etc/postfix/slack_notice.json"
    ```
6. Update alias
    ```
    # newaliases
    ```
