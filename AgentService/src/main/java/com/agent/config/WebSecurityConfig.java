package com.agent.config;

import com.agent.security.TokenAuthenticationFilter;
import com.agent.security.TokenUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.HttpMethod;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.config.annotation.authentication.builders.AuthenticationManagerBuilder;
import org.springframework.security.config.annotation.method.configuration.EnableGlobalMethodSecurity;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.builders.WebSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;
import org.springframework.security.config.http.SessionCreationPolicy;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.security.web.authentication.www.BasicAuthenticationFilter;

@Configuration
@EnableWebSecurity
@EnableGlobalMethodSecurity(prePostEnabled = true)
public class WebSecurityConfig extends WebSecurityConfigurerAdapter {

    @Bean
    public PasswordEncoder passwordEncoder() {
        return new BCryptPasswordEncoder();
    }

    @Autowired
    private CustomUserDetailsService jwtUserDetailsService;

    @Bean
    @Override
    public AuthenticationManager authenticationManagerBean() throws Exception {
        return super.authenticationManagerBean();
    }

    @Autowired
    public void configureGlobal(AuthenticationManagerBuilder auth) throws Exception {
        auth.userDetailsService(jwtUserDetailsService).passwordEncoder(passwordEncoder());
    }

    @Autowired
    private TokenUtil tokenUtils;

    @Override
    protected void configure(HttpSecurity http) throws Exception {
        http
                .sessionManagement().sessionCreationPolicy(SessionCreationPolicy.STATELESS).and()

                .authorizeRequests()
                .antMatchers("/api/users/registerUser").permitAll()
                .antMatchers("/api/users/forgottenpassword").permitAll()
                .antMatchers("/api/users/checkActivationCode").permitAll()
                .antMatchers("/api/users/checkForgottenPassword").permitAll( )
                .antMatchers("/api/users/loginCode").permitAll()
                .antMatchers("api/companies/registerCompany").hasAnyAuthority("User", "Company owner")
                .antMatchers("api/companies/allPendingCompanies").hasAuthority("Admin")
                .antMatchers("api/companies/allApprovedCompanies").hasAnyAuthority("User", "Company owner")
                .antMatchers("api/companies/approveCompanyRequest").hasAuthority("Admin")
                .antMatchers("api/companies/declineCompanyRequest").hasAuthority("Admin")
                .antMatchers("api/companies/myCompanies").hasAnyAuthority("User", "Company owner")
                .antMatchers("api/companies/editCompanyInfo").hasAuthority("Company owner")
                .antMatchers("api/jobs/addOffer").hasAuthority("Company owner")
                .antMatchers("api/jobs/comment").hasAuthority("User")
                .antMatchers("api/jobs/comments/{id}").hasAuthority("User")
                .antMatchers("api/jobs/offers").hasAuthority("User")
                .antMatchers("api/jobs/addSalary").hasAuthority("Company owner")
                .antMatchers("api/jobs/addSurvey").hasAuthority("User")
                .antMatchers("api/jobs/surveys/{id}").hasAuthority("User")


                .anyRequest().authenticated().and()


                .logout().permitAll().and()

                .cors().and()

                .httpBasic().and()

                .addFilterBefore(new TokenAuthenticationFilter(tokenUtils, jwtUserDetailsService), BasicAuthenticationFilter.class);
        http.csrf().disable();
    }

    @Override
    public void configure(WebSecurity web) throws Exception {
        web.ignoring().antMatchers(HttpMethod.POST, "/api/users/login", "api/users/registerUser");
        web.ignoring().antMatchers(HttpMethod.OPTIONS, "/**");
    }
}
