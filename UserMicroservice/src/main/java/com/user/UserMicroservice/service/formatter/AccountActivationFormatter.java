package com.user.UserMicroservice.service.formatter;

public class AccountActivationFormatter {
	public String getText(String link, String siteUrl) {
		String activateURL = siteUrl + "/activation/" + link;
        return " Click on this activation code to verify your registration  " + activateURL;
    }

    public String getSubject() {
        return "Verify your registration at Dislinkt!";
    }

}
