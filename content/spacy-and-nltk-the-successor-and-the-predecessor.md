+++
title = "SpaCy and NLTK: the Successor and the Predecessor"
date = 2019-08-23
+++
[SpaCy](https://spacy.io/) has taken the NLP world by storm. It's reasonable to say that SpaCy is NLTK's successor. Or is it? In this article I will talk about SpaCy, NLTK, and what each is capable of.

# NLTK: the Predecessor

Before we talk about SpaCy, we should talk about NLTK. NLTK was the de facto library to use for NLP, as it had all the features needed to create an NLP model from scratch. Quoting the [website](https://www.nltk.org/):

> NLTK is a leading platform for building Python programs to work with human language data. It provides easy-to-use interfaces to over 50 corpora and lexical resources such as WordNet, along with a suite of text processing libraries for classification, tokenization, stemming, tagging, parsing, and semantic reasoning, wrappers for industrial-strength NLP libraries, and an active discussion forum.

One of my favorite things about NLTK is just how simple it is to access text corpora. It comes with ready-to-use NLP pipeline components like Tokenizers, Part-of-Speech taggers, and Named Entity Recognition. The library has made a name for itself, and has been around for quite the while. Despite all the great features it offers, the problem was that NLTK is based around statistical models. In an era where neural networks are the reigning champions, NLTK wasn't going to cut it.

# SpaCy: the Successor

SpaCy was built to make it easy to train your own models using an intuitive and simple NLP pipeline. The first version of SpaCy was similar to NLTK in that it used statistical models just like NLTK. However, the second version of SpaCy, SpaCy 2.0.0, completely renovated the library. I won't list the feature set here, but the full list is available at the [website](https://spacy.io/usage/v2). The most important aspect of SpaCy 2.0.0 are is the improvements around defining, training and loading your own models and components. Another important new feature is the introduction of CNN models for tagging, parsing and entity recognition. Personally, the best part of SpaCy is the incredibly detailed and friendly documentation. I love just how easy it is to look up anything you need on the website. All of these features came together to do what NLTK couldn't: build powerful, customized NLP pipelines without hassle. 

# What are they for?

It is important to recognize that NLTK and SpaCy is only one part of the NLP ecosystem. The purpose of the libraries is to enable the creation, training and evaluation of models. They make it simple to prepare, preprocess and process text so that scientists and programmers don't need to waste time creating the groundwork to train their models. NLTK was the initial solution to the problem, and with the release of SpaCy into the NLP wilderness, it became the new de facto standard. NLTK is built for statistical models, while SpaCy integrates seamlessly with Tensorflow, PyTorch, and other deep learning frameworks so that it's easier to train your neural networks.

# Which should you use?

If you are planning to train a model to use in an app or service, you should definitely use SpaCy. It's fast, simple, and powerful. There's really only one reason to use NLTK: **to learn**. NLTK, although lacking in certain aspects, is a powerful library, and is paramount to those that are learning about the NLP ecosystem. Since SpaCy is a fairly recent addition to the NLP space, many books that you'll find online will use NLTK. NLTK introduces the fundamentals to you, as well as the problems NLP is trying to solve. You'll get a grasp on the different components of an NLP pipeline, how to process text and how to train models. After getting comfortable with the different components of an NLP pipeline, then you can move to SpaCy. SpaCy also offers a great tutorial [here](https://course.spacy.io/), which was incredibly easy to follow. For any other time, I would use SpaCy. SpaCy objects like as **Doc**, **Token**, and **Span** makes it easy to mutate and annotate text. SpaCy's NLP object allows you to add custom components in a specific order to the NLP pipeline, and can also disable any component, including the ones that come shipped with SpaCy. The library also offers powerful rule-based matchers like Matcher and PhraseMatcher to search for specific text without messing with regex. The image shown below shows the architecture of SpaCy, showing just how simple SpaCy is. 

![SpaCy Architecture Diagram, found at https://spacy.io/architecture-bcdfffe5c0b9f221a2f6607f96ca0e4a.svg](/static/architecture-bcdfffe5c0b9f221a2f6607f96ca0e4a.svg "SpaCy Architecture Diagram")

There are so many reasons to choose SpaCy over NLTK, but you can see for yourself just by following the tutorial.

# Conclusion

We talked about the differences between SpaCy and NLTK, what they are for, and when we should use them. Although they are two different libraries, it doesn't mean they have to be used separately. You can combine NLTK's text corpora with SpaCy's NLP pipeline to make a powerful pipeline with access to extensively annotated texts. SpaCy took over the role of training and processing text, but the other features of NLTK are still relevant and powerful enough to be used in conjunction with SpaCy.
