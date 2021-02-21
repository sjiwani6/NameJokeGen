# NameJokeGen

## About
This program is a web service written in Go hosted on <http://localhost:5000> when it is ran. It receives a random name which then gets inserted into a Chuck Norris style joke. This web service uses two APIs. It uses a random name generator web service and a random Chuck Norris joke generator. The links to the API are below: 
<ul>
<li>Random name generator: https://names.mcquay.me/api/v0</li>
<li>Joke generator: http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=[nerdy]</li>
</ul>

## Instructions to Run
If you don't have Go installed on your computer, run the following command on your terminal to run the binary executable file:
<pre><code>./namejokegen</pre></code>

If you do have Go installed on your computer, run the following command on your terminal:
<pre><code>go run .</pre></code>

Once it has started running, either use your favorite browser to navigate to <code>http://localhost:5000</code> or run the following command on your terminal:
<pre><code>curl 'http://localhost:5000'</pre></code> 