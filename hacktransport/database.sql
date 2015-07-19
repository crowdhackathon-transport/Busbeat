
CREATE TABLE IF NOT EXISTS `Categories` (
  `id` int(11) ,
  `name` varchar(25) ,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `Categories` (`id`, `name`) VALUES
(1, 'music'),
(2, 'culture'),
(3, 'food'),
(4, 'technology'),
(5, 'retails');

CREATE TABLE IF NOT EXISTS `Company` (
  `id` int(11) ,
  `email` varchar(55) ,
  `password` varchar(55) ,
  `name` varchar(255) ,
  `address` varchar(255) ,
  `type` varchar(255) ,
  `webpage` varchar(255) ,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `Company` (`id`, `email`, `password`, `name`, `address`, `type`, `webpage`) VALUES
(1, 'google@google.com', '12345', 'Google', 'silicon valley', '', 'www.google.com');

CREATE TABLE IF NOT EXISTS `Event photos` (
  `id` int(11) ,
  `photo_url` varchar(255) ,
  `description` varchar(555) ,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `Events` (
  `id` int(11) ,
  `name` varchar(255) ,
  `location_long` int(255) ,
  `location_lat` int(255) ,
  `category` int(11) ,
  `photo_URL` varchar(255) ,
  `description` text CHARACTER SET utf8 ,
  `photos_resources` int(11) ,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `Events` (`id`, `name`, `location_long`, `location_lat`, `category`, `photo_URL`, `description`, `photos_resources`) VALUES
(1, 'Hackathon', 24, 38, 0, 'http://crowdhackathon.com/transport/wp-content/uploads/2014/10/Screen-Shot-2015-06-23-at-1.25.06-PM.png', 'Στις 18 και 19 Ιουλίου 2015 οι Συγκοινωνίες Αθηνών (Όμιλος ΟΑΣΑ), η εταιρία καινοτομίας Crowdpolicy και ο Κόμβος Καινοτομίας και Επιχειρηματικότητας της Τεχνόπολις του Δήμου Αθηναίων – INNOVATHENS οργανώνουν διήμερο Crowdhackathon με θέμα την υλοποίηση εφαρμογών σε τομείς της οικονομίας και της κοινωνίας!', 0);

CREATE TABLE IF NOT EXISTS `User` (
  `id` int(11) ,
  `name` varchar(25) ,
  `surname` varchar(25) ,
  `birth year` date ,
  `gender` varchar(6) ,
  `fb_token` int(11) DEFAULT NULL,
  `username` varchar(25) ,
  `password` varchar(255) ,
  `email` varchar(55) ,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `User` (`id`, `name`, `surname`, `birth year`, `gender`, `fb_token`, `username`, `password`, `email`) VALUES
(1, 'Alex', 'Kantas', '1990-09-20', 'male', 0, 'alexzzzboom', 'alex123', 'alexininternet@gmail.com');

CREATE TABLE IF NOT EXISTS `User_likes` (
  `user_id` int(11) ,
  `event_id` int(11) 
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `User_likes` (`user_id`, `event_id`) VALUES
(1, 1);

