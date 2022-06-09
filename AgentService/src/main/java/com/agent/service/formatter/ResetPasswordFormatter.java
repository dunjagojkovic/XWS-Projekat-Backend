package com.agent.service.formatter;

public class ResetPasswordFormatter {
    public String getText(String link, String siteUrl) {
        String resetLink = siteUrl +"/reset/" + link;
        return " Click on this code to reset your password  " + resetLink;
    }

    public String getSubject() {
        return "Reset you password";
    }
}
