package main

import "testing"

func Test1HyphenTrue(t *testing.T) {
	tests := []string{
		"welcome",
		"welcome welcome",
		"welcome welcome welcome",
		"welcome welcome - welcome",
		"welcome, welcome - welcome",
		"welcome, welcome - welcome.",
		"welcome, welcome-welcome", // it's two different words
		"welcome,welcome-welcome",
		"welcome,welcome-welcome",
		"welcome-welcome-welcome", // it's one word
		"-welcome-welcome-welcome-",
		"-wel-come welcome-welcome-", // welcome and welcomewelcome

		"-welcome welcome",
		"-welcome- welcome",
		"welcome- welcome",
		"-welcome -welcome",
		"-welcome- -welcome-",
		"welcome- welcome-",
		"welcome -welcome",
		"welcome- -welcome-",
		"welcome welcome-",
	}

	for _, v := range tests {
		r := Check(&v)
		l := len(r)
		if l != 0 {
			t.Error("expected", 0, "got", l)
		}
	}
}

func Test1HyphenNewlineTrue(t *testing.T) {
	tests := []string{
		"welcome\n",
		"\nwelcome\n",
		"\n\nwelcome\n",
		"\n\n\nwelcome\n",
		"welcome welcome\n",
		"\nwelcome welcome\n",
		"\n\nwelcome welcome\n",
		"\n\n\nwelcome welcome\n",
		"\n\n\nwelcome\nwelcome\n",
		"\n\n\nwelcome\n\nwelcome\n",
		"\n\n\nwelcome\n\n\nwelcome\n",
		"\n\n\nwelcome \n\n\nwelcome\n",
		"\n\n\nwelcome \n\n\n welcome\n",
		"\n\n\nwelcome \n\n\n welcome\n\n",
		"\n\n\nwelcome \n\n\n welcome\n\n\n",
		"welcome\nwelcome welcome",
		"welcome welcome\nwelcome",
		"welcome welcome welcome\n",
		"welcome welcome welcome\n\n",
		"welcome welcome welcome\n\n\n",
		"welcome welcome \nwelcome",
		"welcome welcome \n\nwelcome",
		"welcome welcome \n\n\nwelcome",
		"welcome welcome \n\n\n welcome",
		"welcome welcome\n\n\n welcome",
		"welcome welcome\n welcome",
		"welcome welcome\n\n welcome",
		"welcome welcome\n\n\n welcome",
		"welcome welcome\n\n\nwelcome",
		"welcome welcome\n\nwelcome",
		"welcome welcome\nwelcome",
		"welcome\nwelcome welcome",
		"welcome\n\nwelcome welcome",
		"welcome\n\n\nwelcome welcome",
		"welcome\n\n\n\nwelcome welcome",
		"welcome\n\n\n\n welcome welcome",
		"welcome \n\n\n\n welcome welcome",
		"welcome \n\n\n\nwelcome welcome",
		"welcome \nwelcome welcome",
		"\nwelcome welcome - welcome",
		"welcome\nwelcome - welcome",
		"welcome welcome\n- welcome",
		"welcome welcome -\nwelcome",
		"welcome welcome - welcome\n",
		"welcome welcome - welcome\n",
		"\nwelcome, welcome - welcome",
		"welcome\n, welcome - welcome",
		"welcome,\nwelcome - welcome",
		"welcome, welcome\n- welcome",
		"welcome, welcome -\nwelcome",
		"welcome, welcome - welcome\n",
		"welcome, welcome\n-\nwelcome",
		"\nwelcome, welcome - welcome.",
		"welcome\n, welcome - welcome.",
		"welcome, \nwelcome - welcome.",
		"welcome, welcome\n- welcome.",
		"welcome, welcome -\nwelcome.",
		"welcome, welcome - welcome\n.",
		"welcome, welcome - welcome.\n",
		"welcome, welcome\n-\nwelcome\n.\n",
		"\nwelcome\n,\nwelcome\n-\nwelcome\n.\n",
		"\nwelcome, welcome-welcome",               // it's two different words
		"welcome\n, welcome-welcome",               // it's two different words
		"welcome,\nwelcome-welcome",                // it's two different words
		"welcome, welcome\n-welcome",               // it's two different words
		"welcome, welcome-\nwelcome",               // it's two different words
		"welcome, welcome-welcome\n",               // it's two different words
		"welcome, welcome\n-\nwelcome",             // it's two different words
		"welcome, \nwelcome-welcome",               // it's two different words
		"wel\ncome, welcome-welcome",               // it's two different words
		"welcome, wel\ncome-welcome",               // it's two different words
		"welcome, welcome-wel\ncome",               // it's two different words
		"welcome, w\nelcome-w\ne\nl\nc\no\nm\ne\n", // it's two different words
		"\nwelcome,welcome-welcome",
		"welcome\n,welcome-welcome",
		"welcome,\nwelcome-welcome",
		"welcome,welcome\n-welcome",
		"welcome,welcome-\nwelcome",
		"welcome,welcome-welcome",
		"welcome,welcome-welcome\n",
		"welcome\n-\nwelcome\n-\nwelcome",
		"\n-welcome-welcome-welcome-",
		"-welcome\n-welcome-welcome-",
		"-welcome-\nwelcome-welcome-",
		"-welcome-welcome\n-welcome-",
		"-welcome-welcome-\nwelcome-",
		"-welcome-welcome-welcome\n-",
		"-welcome-welcome-welcome-\n",
		"\n-\n\n\nwelcome\n-welcome\n-welcome-\n",
		"\n-wel-come welcome-welcome-",      // welcome and welcomewelcome
		"-wel\n-come welcome-welcome-",      // welcome and welcomewelcome
		"-wel-\ncome welcome-welcome-",      // welcome and welcomewelcome
		"-wel-come\nwelcome-welcome-",       // welcome and welcomewelcome
		"-wel-come welcome-welcome\n-",      // welcome and welcomewelcome
		"-wel-come\nwelcome-welcome\n-",     // welcome and welcomewelcome
		"-\n\nwel-come\nwelcome-welcome-\n", // welcome and welcomewelcome

		"\n-welcome welcome",
		"-\nwelcome welcome",
		"-\nwelcome- welcome",
		"-welcome\n- welcome",
		"welcome\n- welcome",
		"welcome-\n welcome",
		"-welcome -welcome",
		"\n-welcome\n -welcome",
		"wel\ncome welcome-",
	}

	for _, v := range tests {
		r := Check(&v)
		l := len(r)
		if l != 0 {
			t.Error("expected", 0, "got", l)
		}
	}
}

func TestHyphenFalse(t *testing.T) {
	tests := []string{
		"wel-come welcome",
		"welcome welco-me",
		"w-elcome welcome",
		"we-lcome welcome",
		"welc-ome welcome",
		"welco-me welcome",
		"welcom-e welcome",
		"welcom-e welcome",
		"wel-come people welcome",
		"welcome people welco-me",
		"w-elcome people welcome",
		"we-lcome people welcome",
		"welc-ome people welcome",
		"welco-me people welcome",
		"welcom-e people welcome",
		"welcom-e people welcome",
		"wel-come pe-ople welcome",
		"welcome pe-ople welco-me",
		"w-elcome pe-ople welcome",
		"we-lcome pe-ople welcome",
		"welc-ome pe-ople welcome",
		"welco-me pe-ople welcome",
		"welcom-e pe-ople welcome",
		"welcom-e pe-ople welcome",
	}

	for _, v := range tests {
		r := Check(&v)
		l := len(r)
		if l != 1 {
			t.Error("expected", 1, "got", l)
		}
	}
}

func TestHyphenNewlineFalse(t *testing.T) {
	tests := []string{
		"\nwel-come welcome",
		"wel-come\nwelcome",
		"wel-come welcome\n",
		"\nwelcome welco-me",
		"welcome\nwelco-me",
		"welcome welco-me\n",
		"\nw-elcome welcome",
		"w-elcome\nwelcome",
		"w-elcome welcome\n",
		"\n\nwe-lcome welcome",
		"we-lcome\n\nwelcome",
		"we-lcome welcome\n\n",
		"\nwelc-ome\nwelcome\n",
		"\nwelc-ome welcome\n",
		"\nwelc-ome\nwelcome",
		"\nwelco-me welcome\n",
		"welcom-e\nwelcome\n",
		"\n\n\nwelcom-e\n\n\nwelcome\n\n\n",
		"wel-come people welcome",
		"\nwel-come people welcome",
		"wel-come\npeople welcome",
		"welcome people\nwelco-me",
		"\nw-elcome\npeople welcome",
		"\nwe-lcome people\nwelcome",
		"\nwelc-ome people welcome\n",
		"welco-me\npeople\nwelcome",
		"welcom-e\npeople welcome\n",
		"welcom-e people\nwelcome\n",
		"\nwel-come\npe-ople\nwelcome",
		"\nwelcome\npe-ople welco-me\n",
		"\nw-elcome pe-ople\nwelcome\n",
		"we-lcome\npe-ople\nwelcome\n",
		"welc-ome pe-ople\nwelcome\n",
		"\n\n\nwelco-me\n\n\npe-ople\n\n\nwelcome\n\n\n",
	}

	for _, v := range tests {
		r := Check(&v)
		l := len(r)
		if l != 1 {
			t.Error("expected", 1, "got", l)
		}
	}
}
