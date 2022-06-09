package com.agent.service.formatter;

public class LoginCodeFormatter {
    public String getText(String link, String siteUrl) {
        return " Login on STYLINKT with code: "+link+" \n"+"Hurry up it is valid for 5 minutes!";
    }

    public String getSubject() {
        return "Login by code";
    }
}
